from threading import Semaphore


class Foo:
  def __init__(self):
    self.two = Semaphore(0)
    self.three = Semaphore(0)

  def first(self, printFirst):
    printFirst()
    self.two.release()

  def second(self, printSecond):
    with self.two:
      printSecond()
      self.three.release()

  def third(self, printThird):
    with self.three:
      printThird()
