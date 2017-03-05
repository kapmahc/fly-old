#ifndef FLY_REQUEST_HANDLER_H
#define FLY_REQUEST_HANDLER_H

#include <Poco/Net/HTTPRequestHandler.h>

namespace fly {
class RequestHandler : public Poco::Net::HTTPRequestHandler {
public:
  virtual void handleRequest(Poco::Net::HTTPServerRequest &req,
                             Poco::Net::HTTPServerResponse &resp);
};
}

#endif
