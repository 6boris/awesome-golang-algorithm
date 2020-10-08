import collections
from collections import Callable
from threading import Condition
class H2O:
    def __init__(self):
        self.cv = Condition()
        self.h = collections.deque()
        self.o = collections.deque()

    def hydrogen(self, releaseHydrogen: 'Callable[[], None]') -> None:
        with self.cv:
            # releaseHydrogen() outputs "H". Do not change or remove this line.
            self.h.append(releaseHydrogen)
            self.form()


    def oxygen(self, releaseOxygen: 'Callable[[], None]') -> None:
        with self.cv:
            # releaseOxygen() outputs "O". Do not change or remove this line.
            self.o.append(releaseOxygen)
            self.form()

    def form(self):
        if len(self.h) > 1 and len(self.o) > 0:
            self.h.popleft()()
            self.h.popleft()()
            self.o.popleft()()
        