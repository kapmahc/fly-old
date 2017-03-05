#include "Application.h"
#include "RequestHandler.h"
#include "RequestHandlerFactory.h"
#include <Poco/Net/HTTPServer.h>
#include <iostream>

int fly::Application::main(const std::vector<std::string> &) {
  Poco::Net::HTTPServer s(new fly::RequestHandlerFactory,
                          Poco::Net::ServerSocket(9090),
                          new Poco::Net::HTTPServerParams);

  s.start();
  std::cout << std::endl << "Server started" << std::endl;

  waitForTerminationRequest(); // wait for CTRL-C or kill

  std::cout << std::endl << "Shutting down..." << std::endl;
  s.stop();

  return Poco::Util::Application::EXIT_OK;
}
