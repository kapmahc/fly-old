#ifndef FLY_APPLICATION_H
#define FLY_APPLICATION_H

#include <Poco/Util/ServerApplication.h>
#include <string>
#include <vector>

namespace fly {
class Application : public Poco::Util::ServerApplication {
protected:
  int main(const std::vector<std::string> &);
};
}

#endif
