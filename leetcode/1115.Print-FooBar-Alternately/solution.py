from threading import Semaphore

class FooBar:
    def __init__(self, n):
        self.n = n
        self.foosema = Semaphore(1)
        self.barsema = Semaphore(0)

    def foo(self, printFoo):
        for i in range(self.n):
            self.foosema.acquire()
            printFoo()
            self.barsema.release()

    def bar(self, printBar):
        for i in range(self.n):
            self.barsema.acquire()
            printBar()
            self.foosema.release()