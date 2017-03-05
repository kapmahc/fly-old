#ifndef FLY_APPLICATION_H
#define FLY_APPLICATION_H

#include <Poco/Util/ServerApplication.h>
#include <string>
#include <vector>

namespace fly {

class Application : public Poco::Util::ServerApplication {
public:
  Application();
  ~Application();

protected:
  void initialize(Poco::Util::Application &self);
  void uninitialize();
  void defineOptions(Poco::Util::OptionSet &options);
  void handleHelp(const std::string &name, const std::string &value);
  void handleVersion(const std::string &name, const std::string &value);
  int main(const std::vector<std::string> &);

private:
  bool canRun;
};
}

#endif
