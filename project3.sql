CREATE TABLE Users (
    ID int NOT NULL,
    FirstName varchar(255),
    LastName varchar(255),
    City varchar(255),
    ZipCode int,
    PRIMARY KEY (ID),
    CONSTRAINT FK_UserPassword FOREIGN KEY (UserID)
    REFERENCES UserPasswordHistory(UserID)
    ON DELETE CASCADE
)

CREATE TABLE UserPasswordHistory (
    UserID int NOT NULL,
    Pwd password,
    ChangeDate date,
    CurrentlyActiv boolean
)