describe("Main navigation", function () {
  it("Given I'm on the home page", function () {
    cy.visit("/");
  });

  it("I can visit the reading list page", function () {
    cy.get("a").contains("Reading List").click();
    cy.url().should('include', '/reading-list')
  });

  it("I can visit the favorites page", function () {
    cy.get("a").contains("Favorites").click();
    cy.url().should('include', '/favorites')
  });

  it("I can visit the RSS page", function () {
    cy.get("a").contains("RSS").click();
    cy.url().should('include', '/syndication')
  });

  it("I can visit the account page", function () {
    cy.get("a").contains("Account").click();
    cy.url().should('include', '/account')
  });

  it("and I can go back to the home page", function () {
    cy.get("a").contains("News").click();
    cy.url().should('include', '/')
  });
});
