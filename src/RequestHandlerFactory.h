#ifndef FLY_REQUEST_HANDLER_FACTORY_H
#define FLY_REQUEST_HANDLER_FACTORY_H

#include <Poco/Net/HTTPRequestHandlerFactory.h>

namespace fly {
class RequestHandlerFactory : public Poco::Net::HTTPRequestHandlerFactory {
public:
  virtual Poco::Net::HTTPRequestHandler *
  createRequestHandler(const Poco::Net::HTTPServerRequest &);
};
}

#endif
