Feature: Greeting someone
    Scenario: Time to greet someone
        Given I want to greet <name>
        When I receive a greeting request for <name>
        Then I should greet saying <greeting>
    
        Examples: 
        | name | greeting |
        | d4n13l 4lf4 | Hello d4n13l 4lf4! |
        | Saul | Hello Saul! |
