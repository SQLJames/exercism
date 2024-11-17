"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

EXPECTED_BAKE_TIME = 40
LAYER_PREP_TIME = 2

def bake_time_remaining(minutes_in_oven: int) -> int:
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    return EXPECTED_BAKE_TIME - minutes_in_oven



def preparation_time_in_minutes(layers_of_lasagna: int) -> int:
    """
    Return total amount of prep time needed to prepare the dish.

    This function takes the number of layers and multiplies it by the average prep time per layer.
    """
    return layers_of_lasagna * LAYER_PREP_TIME

def elapsed_time_in_minutes(layers_of_lasagna: int, elapsed_bake_time: int) -> int:
    """
    Return elapsed cooking time.

    This function takes two numbers representing the number of layers & the time already spent
    baking and calculates the total elapsed minutes spent cooking the lasagna.
    """
    prep = preparation_time_in_minutes(layers_of_lasagna)
    return prep + elapsed_bake_time
