import {Given, When, Then} from "cypress-cucumber-preprocessor/steps";
const generator = require('generate-password');
// cy.log(JSON.stringify(csr))

const basic = 'api/v1/user';
const headers = {
    'content-type': 'multipart/form-data',
    'accept': 'application/json'
}
beforeEach(function generateUser() {
    login = generator.generate({})
    email = login + '@gmail.com'
    password  = generator.generate({
        numbers: true,
        symbols: true
    })
});

let user
let login
let email
let password

Then(/^I got response status 201$/,  function () {
    expect(201).to.eq(user.status)
});

Then(/^I got response status 409$/, function () {
    expect(409).to.eq(user.status)
});

Then(/^I got response status 422$/, function () {
    expect(422).to.eq(user.status)
});

Then(/^Error Required Username$/, function () {
    expect("{\"error\":\"Required Username\"}\n").to.eq(user.body)
});

Then(/^Error Required Email$/, function () {
    expect("{\"error\":\"Required Email\"}\n").to.eq(user.body)
});

Then(/^Error Required Password$/, function () {
    expect("{\"error\":\"Required Password\"}\n").to.eq(user.body)
});

Then(/^Error Invalid Email$/, function () {
    expect("{\"error\":\"Invalid Email\"}\n").to.eq(user.body)
});

Then(/^There is no token$/, function () {
    expect("").to.eq(user.body.token)
});

Given(/^I send request for "POST" user$/, function(){
    const csr = cy.fixture('csr.txt').then(() => {
        cy.request({
            method: 'POST',
            url: basic,
            headers: headers,
            form: true,
            body: {
                "login": login,
                "email": email,
                "password": password,
                // TODO add CSR to all steps
                "CSR": csr,
            },
        }).then((resp) => {
            expect(resp.statusText).to.eq('Created')
            user = resp
        })
    })
    // const csr =  cy.readFile('cypress/fixtures/csr.txt');
});

Given(/^I send request for POST user without login$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": "",
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
        cy.log(resp.body)
    })
});

Given(/^I send request for POST user without password$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": email,
            "password": ""
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user without csr$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": password,
            "password": password
            // TODO: ADD CSR.txt
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user without email$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": "",
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        cy.log(resp)
        user = resp
    })
});

Given(/^I send a request for "POST" user twice$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: {
            'content-type': 'multipart/form-data',
            'accept': 'application/json'
        },
        form: true,
        body: {
            "login": login,
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        expect(resp.status).to.eq(409)
        cy.log(resp)
        user = resp
    })
});

Given(/^I send request for POST user with login in field email$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": login,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        cy.log(resp)
        user = resp
    })
});

Given(/^I send request for POST user with email in field login$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": email,
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        cy.log(resp)
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 2 uppercase letters$/, function () {
    let name = generator.generate({
        length: 2,
        lowercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 2 lowercase letters$/, function () {
    let name = generator.generate({
        length: 2,
        uppercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 20 uppercase letters$/, function () {
    let name = generator.generate({
        length: 20,
        lowercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 20 lowercase letters$/, function () {
    let name = generator.generate({
        length: 20,
        uppercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 3 uppercase letters$/, function () {
    let name = generator.generate({
        length: 3,
        lowercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 3 lowercase letters$/, function () {
    let name = generator.generate({
        length: 3,
        uppercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 19 uppercase letters$/, function () {
    let name = generator.generate({
        length: 19,
        lowercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 19 lowercase letters$/, function () {
    let name = generator.generate({
        length: 19,
        uppercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain only numbers$/, function () {
    let name = generator.generate({
        numbers: true,
        uppercase: false,
        lowercase: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain letters in uppercase, lowercase and number$/, function () {
    let name = generator.generate({
        numbers: true
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 2 words with uppercase and lowercase$/, function () {
    let name = generator.generate({
        length: 5,
        symbols: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name + "  " + name,
            "email": email,
            "password": password
        },
    }).then((resp) => {
        expect(resp.statusText).to.eq('Created')
        user = resp
    })
});

Given(/^I send request for POST user with username that contain only 1 letter$/, function () {
    let name = generator.generate({
        length: 1,
        symbols: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with username that contain 21 characters$/, function () {
    let name = generator.generate({
        length: 21,
        symbols: false
    })
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": name,
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with username that contain only spaces$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": "               ",
            "email": email,
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with email that contain 2 @@$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": login + "@@gmail.com",
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with email that not contain domain name$/, function () {
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": login + "@gmail",
            "password": password
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with password that contain 101 characters$/, function () {
    let passw = generator.generate({
        length: 101,
        numbers: true,
        symbols: true,
    })
    cy.log(passw)
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": email,
            "password": passw
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});

Given(/^I send request for POST user with password that contain 100 characters$/, function () {
    let passw = generator.generate({
        length: 100,
        numbers: true,
        symbols: true,
    })
    cy.log(passw)
    cy.request({
        method: 'POST',
        url: basic,
        headers: headers,
        form: true,
        body: {
            "login": login,
            "email": email,
            "password": passw
        },
        failOnStatusCode: false
    }).then((resp) => {
        user = resp
    })
});
