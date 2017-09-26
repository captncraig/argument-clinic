-- Up

CREATE TABLE Pages (
  id INTEGER PRIMARY KEY,
  url TEXT NOT NULL,
  locked BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE Moderators (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  passwordHash TEXT NOT NULL,
  wantsEmails BOOLEAN NOT NULL DEFAULT false,
  isAdmin BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE Comments (
    id INTEGER PRIMARY KEY, 
    pageId INTEGER, 
    raw TEXT,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME NULL,
    needsModeration BOOLEAN NOT NULL DEFAULT false,
    flaggedSpam     BOOLEAN NOT NULL DEFAULT false,
    approvedBy INTEGER NULL,
    approvedOn DATETIME NULL,
    CONSTRAINT Comments_fk_PageId FOREIGN KEY (pageId)
      REFERENCES Pages (id) ON DELETE CASCADE,
    CONSTRAINT Comments_fk_ApproverId FOREIGN KEY (approvedBy)
      REFERENCES Moderators (id) ON UPDATE SET NULL ON DELETE SET NULL
);

CREATE TABLE ModeratorTokens (
  token TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL,
  expires DATETIME NOT NULL
);

CREATE TABLE Settings (
  id INTEGER PRIMARY KEY CHECK (id = 0),
  hasInitialized BOOLEAN NOT NULL DEFAULT 0,
  allowedDomains TEXT NOT NULL DEFAULT '',
  requireModeration BOOLEAN NOT NULL DEFAULT 0,
  checkUrls BOOLEAN NOT NULL DEFAULT 0,

  smtpHost TEXT NOT NULL DEFAULT '',
  smtpUser TEXT NOT NULL DEFAULT '',
  smtpPass TEXT NOT NULL DEFAULT '',
  smtpPort INTEGER NOT NULL DEFAULT 0,
  smtpFrom TEXT NOT NULL DEFAULT '',
  smptpSecure BOOLEAN NOT NULL DEFAULT 0,

  backupRepo TEXT NOT NULL DEFAULT '',
  backupToken TEXT NOT NULL DEFAULT ''
);
INSERT INTO Settings(id) VALUES(0);

-- Down

