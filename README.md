Argument Clinic
===============

Argument clinic is a simple comment server written in node. It has a few primary goals:

1. **Easy** to self-host. Designed to run for free on [glitch](glitch.com)
2. Privacy Focused. **Absolutely no tracking.**
3. Simple open-api. Doesn't require magic js include.
4. You own the data. Sqlite backed, but easy to import/export.

Comment Features
-----

- Any user can comment immediately. No social logins required (or possible).
- Threaded comments and replys.
- Optional email notification of replies.
- Edit and deletion by original commenter only.
- Configurable markdown capabilities.
- Can put comments on any page, or even a section of a page.

Non-features
-----

- Social logins, or any kind of "log in to post". 
- Up/Down voting.

Moderation Features
-----

- You may choose to require a **moderation queue** for all comments.
- May have more than one moderator for a site.
- Askimet integration for automated spam detection. Seperate spam queue.
- Email moderators on all comments.
- All settings editable in simple web ui.