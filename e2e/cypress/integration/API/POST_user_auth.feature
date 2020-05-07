@api
@post
# ./node_modules/.bin/cypress-tags run -e TAGS='@'

Feature: POST
  Authentication for users to get in to the system and receive JWT token
  Login user into the system

  Background:
    Given Generate data of user
    When I send request for POST user
    Then  I got response status 201
#  username
#  email
#  password
#  csr.txt

  @positive
  Scenario: (POST) login
    Given I send request for login user
    When I got response status 201
    And Description "Certificate"
    Then Response body is empty

  @negative
  Scenario: (POST) user cannot register again
    Given I send request for POST user
    And I got response status 201
    And Description "Certificate"
    When I send request for POST user
    Then I got response status 400
    And Description "User already exist"



