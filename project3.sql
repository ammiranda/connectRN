CREATE TABLE Users (
    ID INT NOT NULL,
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    City VARCHAR(255),
    ZipCode INT,
    PRIMARY KEY (ID),
    CONSTRAINT FK_UserPassword FOREIGN KEY (UserID)
    REFERENCES UserPasswordHistory(UserID)
    ON DELETE CASCADE
)

CREATE TABLE UserPasswordHistory (
    UserID INT NOT NULL,
    Pwd VARCHAR(255),
    ChangeDate DATETIME
    CONSTRAINT CD_changeDate DEFAULT (getdate)),
    CurrentlyActive BOOLEAN DEFAULT TRUE,
)

SELECT Pwd FROM UserPasswordHistory WHERE CurrentlyActive = TRUE;

INSERT INTO UserPasswordHistory (Pwd) VALUES (HASHBYTES('MD5', 'Password123'));

