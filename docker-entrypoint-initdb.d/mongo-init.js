db.auth('admin', 'admin')

db = db.getSiblingDB('Metadata')

db.createUser({
    user: 'user',
    pwd: 'user',
    roles: [
        {
            role: 'root',
            db: 'Metadata',
        },
    ],
});