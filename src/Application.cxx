#include "Application.h"
#include "Configuration.h"
#include "RequestHandler.h"
#include "RequestHandlerFactory.h"
#include <Poco/Logger.h>
#include <Poco/Message.h>
#include <Poco/Net/HTTPServer.h>
#include <iostream>

int fly::Application::main(const std::vector<std::string> &) {
  Poco::Logger &logger = Poco::Logger::get("fly");
  logger.setLevel(Poco::Message::PRIO_TRACE);

  Poco::Net::HTTPServer s(new fly::RequestHandlerFactory,
                          Poco::Net::ServerSocket(9090),
                          new Poco::Net::HTTPServerParams);

  s.start();
  logger.information("server started with logging level %d", logger.getLevel());
  // wait for CTRL-C or kill
  waitForTerminationRequest();
  logger.information("shutting down...");
  s.stop();

  return Poco::Util::Application::EXIT_OK;
}
