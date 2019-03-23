package cluster

//go:generate go get github.com/golang/mock/gomock
//go:generate go install github.com/golang/mock/mockgen
//go:generate mockgen -destination=../util/mocks/mock_$GOPACKAGE/hash.go -package=mock_$GOPACKAGE -source hash.go
//go:generate gofmt -s -l -w ../util/mocks/mock_$GOPACKAGE/hash.go
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../util/mocks/mock_$GOPACKAGE/hash.go

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/startup"
	"github.com/openshift/openshift-azure/pkg/sync"
)

type Hasher interface {
	HashScaleSet(*api.OpenShiftManagedCluster, *api.AgentPoolProfile) ([]byte, error)
	HashSyncPod(cs *api.OpenShiftManagedCluster) ([]byte, error)
}

type hasher struct {
	log            *logrus.Entry
	testConfig     api.TestConfig
	startupFactory func(*logrus.Entry, *api.OpenShiftManagedCluster) (startup.Interface, error)
}

// HashScaleSet returns the hash of a scale set
func (h *hasher) HashScaleSet(cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile) ([]byte, error) {
	hash := sha256.New()

	// the hash is invariant of name, suffix, count
	appCopy := *app
	appCopy.Count = 0
	appCopy.Name = ""

	// and also the SAS URIs
	csCopy := cs.DeepCopy()
	csCopy.Config.MasterStartupSASURI = ""
	csCopy.Config.WorkerStartupSASURI = ""

	vmss, err := arm.Vmss(csCopy, &appCopy, "", "", h.testConfig) // TODO: backupBlob is rather a layering violation here
	if err != nil {
		return nil, err
	}

	err = json.NewEncoder(hash).Encode(vmss)
	if err != nil {
		return nil, err
	}

	s, err := h.startupFactory(h.log, cs)
	if err != nil {
		return nil, err
	}

	b, err := s.Hash(app.Role)
	if err != nil {
		return nil, err
	}

	hash.Write(b)

	if app.Role == api.AgentPoolProfileRoleMaster {
		// add certificates pulled from keyvault by the master to the hash, to
		// ensure the masters update if a cert changes.  We don't add the keys
		// because these are not necessarily stable (sometimes the 'D' value of
		// the RSA key returned by keyvault differs to the one that was sent).
		// I believe that in a given RSA key, there are multiple suitable values
		// of 'D', so this is not a problem, however it doesn't make the value
		// suitable for a hash.  References:
		// https://stackoverflow.com/a/14233140,
		// https://crypto.stackexchange.com/a/46572.
		hash.Write(cs.Config.Certificates.OpenShiftConsole.Certs[0].Raw)
		hash.Write(cs.Config.Certificates.Router.Certs[0].Raw)
	}

	return hash.Sum(nil), nil
}

// HashSyncPod returns the hash of the sync pod output
func (h *hasher) HashSyncPod(cs *api.OpenShiftManagedCluster) ([]byte, error) {
	s, err := sync.New(h.log, cs)
	if err != nil {
		return nil, err
	}

	return s.Hash()
}
