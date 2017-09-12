-- Up

CREATE TABLE Pages (id INTEGER PRIMARY KEY, url string);
CREATE TABLE Comments (
    id INTEGER PRIMARY KEY, 
    pageId INTEGER, 
    raw TEXT,
    created DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated DATETIME NULL,
    CONSTRAINT Comments_fk_PageId FOREIGN KEY (pageId)
      REFERENCES Pages (id) ON DELETE CASCASE
);

-- Down

