# Crypter

Crypter encrypts files in the whole directory as bcrypt / argon2id hash. Reads
from the environment

Password can be any string, and then that is used as a SECRET for deplyoment to a environment.
The secret is passed in as an ENV is "BS_CRYPT_PASSWORD".

