
import sys, os, time
from heapq import nlargest

class ExtractHighestNumbers:

    """
    Used for extracting top X highest numbers from a file that contains
    individual numbers on each line.
    """

    def __init__(self, numbers_file=None, highest_number_count=None):
        self.numbers_file = numbers_file
        self.highest_number_count = highest_number_count
        self.all_numbers = self.__extract_numbers_from_file()
        self.top_x_numbers = self.__get_highest_numbers()

    def __timing_val(func):
        """
        Decorator used for measuring time execution of functions.
        """

        def wrapper(*arg, **kw):
            t1 = time.time()
            result = func(*arg, **kw)
            total_time = time.time() - t1
            print(f'{func.__name__} ran in {total_time} seconds.')
            return result
        return wrapper

    @__timing_val
    def __extract_numbers_from_file(self):
        """
        Populate all_numbers list with all the numbers from file.
        """

        numbers = []

        if os.path.isfile(self.numbers_file):
            with open(self.numbers_file) as f:
                for line in f:
                    try:
                        numbers.append(int(line.rstrip()))
                    except:
                        continue
            return list(set(numbers))    # return list with unique numbers
        else:
            raise Exception(f'{self.numbers_file} does not exist!')

    @__timing_val
    def __get_highest_numbers(self):
        """
        Get highest X numbers from list of all numbers.
        """

        # slower 1
        # self.all_numbers.sort(reverse=True)
        # return self.all_numbers[:self.highest_number_count]

        return nlargest(self.highest_number_count, self.all_numbers)
        


if __name__ == "__main__":

    # define vars
    file = sys.argv[1]
    number_of_highest_numbers = int(sys.argv[2])

    # get highest numbers from file
    numbers = ExtractHighestNumbers(file, number_of_highest_numbers)
    for number in numbers.top_x_numbers:
        print(number)


    