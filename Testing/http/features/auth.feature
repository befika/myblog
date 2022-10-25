Feature: Auth api

  In order to use  auth API
  as an API auth
  I need to be able to managge Auth

  Scenario: user should login
    Given the user enters email and password:
      | email                | password |
      | befika77@gmail.com   | 123456   |
    When I send "POST" HTTP request to Auth "/v1/auth/login"
    Then the auth response code should be 200