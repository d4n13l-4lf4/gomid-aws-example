Feature: Authenticated users should be able to use my service

    Scenario: User authenticated succesfully
        Given <name> is authenticated
        When <name> submits a request to my service
        Then I should allow access to my service

        Examples:
        | name |
        | d4n13l 4lf4 |
        | Saul |

    Scenario: User is not allowed to use my service
        Given <name> is not authenticated
        When <name> submits a request to my service
        Then I should deny access saying <message> with <status_code> status code

        Examples:
        | status_code | name | message |
        | 403 | d4n13l 4lf4 | d4n13l 4lf4 is not allowed to access this resource | 
        | 403 | Saul | Saul is not allowed to access this resource | 

    Scenario: User submits a wrong request to my service
        Given <name> is authenticated
        When <name> submits a wrong request to my service
        Then I should deny access saying <message> with <status_code> status code

        Examples:
        | status_code | name | message |
        | 403 | d4n13l 4lf4 | wrong request for authorization |
        | 403 | Saul | wrong request for authorization |