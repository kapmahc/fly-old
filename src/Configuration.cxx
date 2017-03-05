#include "Configuration.h"

std::string filename(std::string name) { return name + ".properties"; }

fly::Configuration::Configuration(std::string name) {
  this->file = filename(name);
  Poco::AutoPtr<Poco::Util::PropertyFileConfiguration> cfg(
      new Poco::Util::PropertyFileConfiguration(this->file));
  this->cfg = cfg;
}

void fly::Configuration::generate(std::string name) {
  std::string fn = filename(name);
  Poco::Util::PropertyFileConfiguration *cfg =
      new Poco::Util::PropertyFileConfiguration(fn);
  cfg->setBool("app.debug", true);
  cfg->setInt("app.port", 8080);
  cfg->setString("app.name", "fly");
  cfg->save(fn);
}

std::string fly::Configuration::name() { return cfg->getString("server.name"); }
int fly::Configuration::port() { return cfg->getInt("server.port"); }
bool fly::Configuration::debug() { return cfg->getBool("server.debug"); }
