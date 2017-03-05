#ifndef FLY_CONFIGURATION_H
#define FLY_CONFIGURATION_H

#include <Poco/Util/PropertyFileConfiguration.h>
#include <string>

namespace fly {
class Configuration {
public:
  Configuration(std::string name);
  std::string name();
  int port();
  bool debug();
  static void generate(std::string name);

private:
  Poco::AutoPtr<Poco::Util::PropertyFileConfiguration> cfg;
  std::string file;
};
}
#endif
