#include "RequestHandler.h"
#include <Poco/Net/HTTPServerRequest.h>
#include <Poco/Net/HTTPServerResponse.h>
#include <iostream>

void fly::RequestHandler::handleRequest(Poco::Net::HTTPServerRequest &req,
                                        Poco::Net::HTTPServerResponse &resp) {
  resp.setStatus(Poco::Net::HTTPResponse::HTTP_OK);
  resp.setContentType("text/html");

  std::ostream &out = resp.send();
  out << "<h1>Hello world!</h1>"
      << "<p>Host: " << req.getHost() << "</p>"
      << "<p>Method: " << req.getMethod() << "</p>"
      << "<p>URI: " << req.getURI() << "</p>";
  out.flush();

  std::cout << " and URI=" << req.getURI() << std::endl;
}
