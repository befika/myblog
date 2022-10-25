Feature: users crud api

  In order to use  users API
  as an API users
  I need to be able to managge users

  Scenario: should get users
    Given there are users:
      | id                                   | username  | password | phone      | first_name  | middle_name | last_name | email              | created_at                       | updated_at                       |
      | cbf77990-a47c-4022-85d6-99cc76bd705a | befika77  | 12345678 | 0987654321 | befikadu    | shumet      | alibew    | befika77@gmail.com | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
      | f318e33c-2918-413b-be4b-09a6d3496e6e | testuser1 | 12345688 | 0987654456 | testuser    | testuser    | testuser  | testuser@gmail.com | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
    When I send "GET" HTTP request to "/v1/users"
    Then the response code should be 200
    
  # Scenario: should create a users
  #   Given  I have request json:
  #     """
  #     {
  #        "username": "clerk123",
  #        "password": "12345",
  #        "phone": "09876543299",
  #        "first_name": clerk,
  #        "middle_name": "testuser",
  #        "last_name": "clerk",
  #        "email": "clerk@gmail.com",
  #        "rolename":"SYSTEM-ADMIN"
  #     }
  #     """
  #   When I send "POST" request to "/v1/users"
  #   Then the response code should be 200

  # Scenario: should update a users
  #   Given there are users:
  #     | id                                   | username  | password | phone      | first_name  | middle_name | last_name | email              | rolename     | created_at                       | updated_at                       |
  #     | cbf77990-a47c-4022-85d6-99cc76bd705a | befika77  | 12345678 | 0987654321 | befikadu    | shumet      | alibew    | befika77@gmail.com | SYSTEM-ADMIN | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
  #     | f318e33c-2918-413b-be4b-09a6d3496e6e | testuser1 | 12345688 | 0987654456 | testuser    | testuser    | testuser  | testuser@gmail.com | SYSTEM-CLERK | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
  #   And I have request json:
  #     """
  #     {
  #       "username": "testuser_update",
  #       "phone": "0951109200",
  #       "first_name": clerk3,
  #       "middle_name": "clerk3",
  #     }
  #     """
  #   When I send "PATCH" request to "/v1/users/cbf77990-a47c-4022-85d6-99cc76bd705a"
  #   Then the response code should be 200

  # Scenario: should delete a user
  #   Given there are users:
  #     | id                                   | first_name | middle_name | last_name | email             | phone_number | username | password  | created_at                       | updated_at                       |
  #     | cbf77990-a47c-4022-85d6-99cc76bd705a | befikadu   | shumet      | alibew    | befika77@gmail.com| 0938904989   | befika77 | 123456    | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
  #     | f318e33c-2918-413b-be4b-09a6d3496e6e | abebe      | kebede      | abebe     | abebe@gmail.com   | 0987878787   | abebe7   | 123456    | 2021-10-27T17:11:07.932352+03:00 | 2021-10-27T17:11:07.932352+03:00 |
  #   When  I send "DELETE" request to "/v1/users/cbf77990-a47c-4022-85d6-99cc76bd705a"
  #   Then the response code should be 200
