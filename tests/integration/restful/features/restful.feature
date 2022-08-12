Feature: Health check

   Scenario: Get successful ping server
      Given nothing
      When I send "GET" request to "/api/ping"
      Then the response code should be 200
      And the response time should be less than 7 seconds
      And the response should match json:
         """
         {
            "url": "/api/ping",
            "date": "2022-08-11T17:41:08.969439651+07:00"
         }
         """
      Examples:
         | url  |
         | date |