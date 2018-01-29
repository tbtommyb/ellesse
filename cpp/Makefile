CXX = clang++
CXXFLAGS = -std=c++1z -Wall
LIBS = -lboost_system -lboost_filesystem
SRCDIR = src
OUTDIR = build

ellesse: main.cpp $(OUTDIR)/ellesse.o
	$(CXX) main.cpp $(OUTDIR)/ellesse.o -o $(OUTDIR)/$@ $(CXXFLAGS) $(LIBS)

$(OUTDIR)/ellesse.o: $(SRCDIR)/ellesse.cpp
	$(CXX) -c $(SRCDIR)/ellesse.cpp -o $@ $(CXXFLAGS)

clean:
	cd $(OUTDIR) && rm ellesse ellesse.o
