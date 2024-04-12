-- +goose Up
-- +goose StatementBegin
INSERT INTO category (category_name) VALUES 
('Bills'),
('Debt'),
('Education'),
('Entertainment'),
('Family'),
('Food & Drinks'),
('Savings'),
('Shopping'),
('Top Up'),
('Transportation'),
('Social Events'),
('Others');

-- Bills Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Credit Card', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Electricity', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Gas', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Insurance', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Internet', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Landline', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Maintenance Fee', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Mobile & Data', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Rent', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Subscriptions', (SELECT category_id FROM category WHERE category_name = 'Bills')),
('Water', (SELECT category_id FROM category WHERE category_name = 'Bills'));

-- Education Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Books', (SELECT category_id FROM category WHERE category_name = 'Education')),
('Online Courses', (SELECT category_id FROM category WHERE category_name = 'Education')),
('Tuition Fee', (SELECT category_id FROM category WHERE category_name = 'Education'));

-- Entertainment Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Concert', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Games', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Hangout', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Hobby', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Movies', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Streaming Services', (SELECT category_id FROM category WHERE category_name = 'Entertainment')),
('Vacation', (SELECT category_id FROM category WHERE category_name = 'Entertainment'));

-- Food & Drinks Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Cafe', (SELECT category_id FROM category WHERE category_name = 'Food & Drinks')),
('Restaurant', (SELECT category_id FROM category WHERE category_name = 'Food & Drinks')),
('Take Outs', (SELECT category_id FROM category WHERE category_name = 'Food & Drinks'));

-- Healthcare Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Medicine', (SELECT category_id FROM category WHERE category_name = 'Healthcare')),
('Gym', (SELECT category_id FROM category WHERE category_name = 'Healthcare')),
('Medical fee', (SELECT category_id FROM category WHERE category_name = 'Healthcare')),
('Sports', (SELECT category_id FROM category WHERE category_name = 'Healthcare'));

-- Savings Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Emergency Fund', (SELECT category_id FROM category WHERE category_name = 'Savings')),
('House', (SELECT category_id FROM category WHERE category_name = 'Savings')),
('Investment', (SELECT category_id FROM category WHERE category_name = 'Savings')),
('Pension', (SELECT category_id FROM category WHERE category_name = 'Savings')),
('Vacation', (SELECT category_id FROM category WHERE category_name = 'Savings'));

-- Shopping Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Fashion', (SELECT category_id FROM category WHERE category_name = 'Shopping')),
('Gadget & Electronics', (SELECT category_id FROM category WHERE category_name = 'Shopping')),
('Groceries', (SELECT category_id FROM category WHERE category_name = 'Shopping'));

-- Social Events Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Charity & Donations', (SELECT category_id FROM category WHERE category_name = 'Social Events')),
('Funeral', (SELECT category_id FROM category WHERE category_name = 'Social Events')),
('Gifts', (SELECT category_id FROM category WHERE category_name = 'Social Events')),
('Wedding', (SELECT category_id FROM category WHERE category_name = 'Social Events'));

-- Transportation Sub-Categories
INSERT INTO category (category_name, parent_category_id) VALUES 
('Gasoline', (SELECT category_id FROM category WHERE category_name = 'Transportation')),
('Parking Fee', (SELECT category_id FROM category WHERE category_name = 'Transportation')),
('Public Transport', (SELECT category_id FROM category WHERE category_name = 'Transportation')),
('Taxi / Ride Hailing', (SELECT category_id FROM category WHERE category_name = 'Transportation')),
('Travel Fares', (SELECT category_id FROM category WHERE category_name = 'Transportation')),
('Vehicle Maintenance', (SELECT category_id FROM category WHERE category_name = 'Transportation'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM category;
-- +goose StatementEnd
