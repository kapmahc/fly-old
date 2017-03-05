#include "RequestHandler.h"
#include "RequestHandlerFactory.h"

Poco::Net::HTTPRequestHandler *fly::RequestHandlerFactory::createRequestHandler(
    const Poco::Net::HTTPServerRequest &) {
  return new fly::RequestHandler;
}
