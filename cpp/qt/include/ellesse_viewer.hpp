#ifndef __ellesse_viewer__
#define __ellesse_viewer__

#include <QObject>
#include "ellesse.hpp"

class EllesseViewer : public QObject
{
  Q_OBJECT
public:
  EllesseViewer();
  EllesseViewer(const std::string&);
  ~EllesseViewer() = default;
  QStringList fileList();
  Q_PROPERTY(QStringList fileList READ fileList NOTIFY fileListChanged)
private:
  std::string path;
  Ellesse directory;
signals:
  void fileListChanged();
public slots:
  void change(const QString&);
};

#endif
