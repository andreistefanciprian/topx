
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
        self._all_numbers = self._extract_numbers_from_file()
        self.top_x_numbers = self._get_highest_numbers()

    def timing_val(func):
        def wrapper(*arg, **kw):
            t1 = time.time()
            result = func(*arg, **kw)
            total_time = time.time() - t1
            print(f'{func.__name__} ran in {total_time} seconds.')
            return result
        return wrapper

    @timing_val
    def _extract_numbers_from_file(self):
        """
        Populate all_numbers list witth all the numbers from file.
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

    @timing_val   
    def _get_highest_numbers(self):
        """
        Get highest X numbers from list of all numbers.
        """

        # slower 1
        # self._all_numbers.sort(reverse=True)
        # return self._all_numbers[:self.highest_number_count]

        return nlargest(self.highest_number_count, self._all_numbers)
        


if __name__ == "__main__":

    # define vars
    file = sys.argv[1]
    number_of_highest_numbers = int(sys.argv[2])

    # get highest numbers from file
    numbers = ExtractHighestNumbers(file, number_of_highest_numbers)
    for number in numbers.top_x_numbers:
        print(number)


    