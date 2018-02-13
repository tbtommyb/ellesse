#include "ellesse_viewer.hpp"

EllesseViewer::EllesseViewer() : EllesseViewer{"."}
{
};

EllesseViewer::EllesseViewer(const std::string& path) :
  path{path},
  directory{Ellesse::Ellesse{path}}
{
};

QStringList EllesseViewer::fileList()
{
  QStringList list{};

  auto items = directory.items();
  for(const auto& f : items) {
    list << QString{f.pathName.c_str()};
  }
  return list;
};

void EllesseViewer::change(const QString& newPath)
{
  path = path + "/" + newPath.toStdString();
  directory = Ellesse::Ellesse{path};
  emit fileListChanged();
};

#include "moc_ellesse_viewer.cpp"
