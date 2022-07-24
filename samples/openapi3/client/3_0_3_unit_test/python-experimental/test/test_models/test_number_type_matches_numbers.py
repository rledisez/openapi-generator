# coding: utf-8

"""
    openapi 3.0.3 sample spec

    sample spec for testing openapi functionality, built from json schema tests for draft6  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by: https://openapi-generator.tech
"""

import unittest

import unit_test_api
from unit_test_api.model.number_type_matches_numbers import NumberTypeMatchesNumbers
from unit_test_api import configuration


class TestNumberTypeMatchesNumbers(unittest.TestCase):
    """NumberTypeMatchesNumbers unit test stubs"""
    _configuration = configuration.Configuration()

    def test_an_array_is_not_a_number_fails(self):
        # an array is not a number
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                [
                ],
                _configuration=self._configuration
            )

    def test_null_is_not_a_number_fails(self):
        # null is not a number
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                None,
                _configuration=self._configuration
            )

    def test_an_object_is_not_a_number_fails(self):
        # an object is not a number
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                {
                },
                _configuration=self._configuration
            )

    def test_a_boolean_is_not_a_number_fails(self):
        # a boolean is not a number
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                True,
                _configuration=self._configuration
            )

    def test_a_float_is_a_number_passes(self):
        # a float is a number
        NumberTypeMatchesNumbers._from_openapi_data(
            1.1,
            _configuration=self._configuration
        )

    def test_a_string_is_still_not_a_number_even_if_it_looks_like_one_fails(self):
        # a string is still not a number, even if it looks like one
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                "1",
                _configuration=self._configuration
            )

    def test_a_string_is_not_a_number_fails(self):
        # a string is not a number
        with self.assertRaises((unit_test_api.ApiValueError, unit_test_api.ApiTypeError)):
            NumberTypeMatchesNumbers._from_openapi_data(
                "foo",
                _configuration=self._configuration
            )

    def test_an_integer_is_a_number_passes(self):
        # an integer is a number
        NumberTypeMatchesNumbers._from_openapi_data(
            1,
            _configuration=self._configuration
        )

    def test_a_float_with_zero_fractional_part_is_a_number_and_an_integer_passes(self):
        # a float with zero fractional part is a number (and an integer)
        NumberTypeMatchesNumbers._from_openapi_data(
            1.0,
            _configuration=self._configuration
        )


if __name__ == '__main__':
    unittest.main()