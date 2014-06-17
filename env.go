package mallory

import (
	"errors"
	"flag"
	"os"
	"path"
)

const (
	CO_RED    = "\033[0;31;49m"
	CO_GREEN  = "\033[0;32;49m"
	CO_YELLOW = "\033[0;33;49m"
	CO_BLUE   = "\033[0;34;49m"
	CO_RESET  = "\033[0m"
)

// Provide global config for mallory
type Env struct {
	// work space, default is $HOME/.mallory
	Work string
	// local addr to listen and serve, default is 127.0.0.1:18087
	Addr string
	// remote engine to be used, "gae" or "direct"(default)
	Engine string
	// GAE application ID, only valid when the engine is "gae"
	// e.g. kill-me-baby of http://kill-me-baby.appspot.com
	AppSpot string
	// url of socks proxy, only valid when the engine is SocksToHttp
	SocksProxy string
	// > http://www.akadia.com/services/ssh_test_certificate.html
	// > http://mitmproxy.org/doc/ssl.html
	// RSA private key file and self-signed root certificate file
	// Can be generated by OpenSSL:
	// - RSA private key file, without input any extra info
	//      openssl genrsa -out mallory.key 2048
	// - Self-signed root certificate file, need input some X.509 attributes
	//   such as Country Name, Comman Name etc.
	//      openssl req -new -x509 -days 365 -key mollory.key -out mallory.crt
	Key  string // mallory.key
	Cert string // mallory.crt
	// pac file path
	PAC string
	// terminal helper, test the default logger(os.Stderr) is terminal or not
	Istty bool
}

// Prepare flags and env
func (self *Env) Parse() error {
	flag.StringVar(&self.Addr, "addr", "127.0.0.1:18087", "Mallory server address, Host:Port")
	// -appsopt=debug to connect the localhost server for debug
	flag.StringVar(&self.AppSpot, "appspot", "oribe-yasuna", "GAE application ID, only valid when engine is gae")
	flag.StringVar(&self.Engine, "engine", "direct", `Mallory engine, "direct", "s2h" or "gae"`)
	flag.StringVar(&self.SocksProxy, "socks_proxy", "socks5://localhost:1314", "SOCKS5 proxy URL, only valid when engine is s2h")
	flag.StringVar(&self.Work, "work", path.Join("$HOME", ".mallory"), "Work directory for mallory")

	flag.Parse()

	if self.Engine != "gae" && self.Engine != "direct" && self.Engine != "s2h" {
		return errors.New(`engine should be "direct", "s2h" or "gae"`)
	}

	// expand env vars for paths
	self.Work = os.ExpandEnv(self.Work)
	self.Key = path.Join(self.Work, "mallory.key")
	self.Cert = path.Join(self.Work, "mallory.crt")
	self.PAC = path.Join(self.Work, "mallory.pac")

	self.Istty = Isatty(os.Stderr)
	return nil
}
