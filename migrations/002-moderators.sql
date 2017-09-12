-- Up

ALTER TABLE Comments ADD needsModeration BOOLEAN NOT NULL DEFAULT false;
ALTER TABLE Comments ADD flaggedSpam     BOOLEAN NOT NULL DEFAULT false;

CREATE TABLE Moderators (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  passwordHash TEXT NOT NULL,
  isAdmin BOOLEAN NOT NULL DEFAULT false
);

ALTER TABLE Comments ADD approvedBy INTEGER NULL;
ALTER TABLE Comments ADD approvedOn DATETIME NULL;
ALTER TABLE Comments ADD CONSTRAINT Comments_fk_ApproverId FOREIGN KEY (approvedBy)
      REFERENCES Moderators (id) ON UPDATE SET NULL ON DELETE SET NULL

CREATE TABLE ModeratorTokens (
  token TEXT,
  email TEXT,
  expires DATETIME NOT NULL,
);


-- Down
