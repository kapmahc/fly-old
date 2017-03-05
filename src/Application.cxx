#include "Application.h"
#include "Configuration.h"
#include "RequestHandler.h"
#include "RequestHandlerFactory.h"
#include <Poco/Logger.h>
#include <Poco/Message.h>
#include <Poco/Net/HTTPServer.h>
#include <Poco/Util/HelpFormatter.h>
#include <iostream>

fly::Application::Application() { canRun = true; }

fly::Application::~Application() {}

void fly::Application::initialize(Poco::Util::Application &self) {
  loadConfiguration();
  Poco::Util::ServerApplication::initialize(self);
}

void fly::Application::uninitialize() { ServerApplication::uninitialize(); }

void fly::Application::defineOptions(Poco::Util::OptionSet &options) {
  Poco::Util::ServerApplication::defineOptions(options);
  options.addOption(Poco::Util::Option("version", "v", "Show version")
                        .required(false)
                        .repeatable(false)
                        .callback(Poco::Util::OptionCallback<Application>(
                            this, &Application::handleVersion)));
  options.addOption(
      Poco::Util::Option("help", "h", "Display argument help information")
          .required(false)
          .repeatable(false)
          .callback(Poco::Util::OptionCallback<Application>(
              this, &Application::handleHelp)));
}

void fly::Application::handleVersion(const std::string &name,
                                     const std::string &value) {
  stopOptionsProcessing();
  std::cout << GIT_VERSION << "(" << BUILD_TIME << ")" << std::endl;
  canRun = false;
}

void fly::Application::handleHelp(const std::string &name,
                                  const std::string &value) {
  Poco::Util::HelpFormatter helpFormatter(options());
  helpFormatter.setCommand(commandName());
  helpFormatter.setUsage("OPTIONS");
  helpFormatter.setHeader("A complete open source ecommerce solution.");
  helpFormatter.format(std::cout);
  stopOptionsProcessing();

  canRun = false;
}

int fly::Application::main(const std::vector<std::string> &) {
  if (canRun) {
    Poco::Logger &logger = Poco::Logger::get("fly");
    logger.setLevel(Poco::Message::PRIO_TRACE);

    Poco::Net::HTTPServer srv(new fly::RequestHandlerFactory,
                              Poco::Net::ServerSocket(9090),
                              new Poco::Net::HTTPServerParams);

    srv.start();
    logger.information("server started with logging level %d",
                       logger.getLevel());
    // wait for CTRL-C or kill
    waitForTerminationRequest();
    logger.information("shutting down...");
    srv.stop();
  }
  return Poco::Util::Application::EXIT_OK;
}
