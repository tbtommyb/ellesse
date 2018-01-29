CCC = clang++

CCCFLAGS = -std=c++1z
LIBS = -lboost_system -lboost_filesystem

ellesse:
	$(CCC) main.cpp src/*.cpp -o ellesse $(CCCFLAGS) $(LIBS)

clean:
	rm ellesse
