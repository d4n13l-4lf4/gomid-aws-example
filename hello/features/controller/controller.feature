Feature: Greeting someone through my greeting API

    Scenario: Greeting successfully
        Given I receive a greeting request for <name>
        When I greet as <name>
        Then I should greet <name> successfully with <status_code>

        Examples:
        | status_code | name |
        | 200 | d4n13l 4lf4 |
        | 200 | Saul |

    Scenario: Invalid greeting request
        Given I receive an invalid greeting request
        When I greet as Hello!
        Then I should get an error <message> with <status_code>

        Examples:
        | status_code | message |
        | 400 | bad request |
    
    Scenario: Throwing an error while greeting
        Given I receive a greeting request for <name>
        When Greeting fails with error <message>
        Then I should get an error <message> with <status_code>
       
       Examples:
        | status_code | name | message |
        | 500 | d4n13l 4lf4 | could not greet d4n13l 4lf4 |
        | 500 | Saul | could not greet Saul |
