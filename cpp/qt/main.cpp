#include <QGuiApplication>
#include <QQmlContext>
#include <qqmlapplicationengine>
#include <iostream>
#include "ellesse_viewer.hpp"

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);

    QCoreApplication::addLibraryPath("./");

    EllesseViewer viewer{};

    QQmlApplicationEngine engine;

    QQmlContext *context = engine.rootContext();
    context->setContextProperty("_Ellesse", &viewer);

    engine.load(QUrl(QStringLiteral("qrc:/main.qml")));
    return app.exec();
}
