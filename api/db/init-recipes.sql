CREATE table recipes (
    id SERIAL PRIMARY KEY,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE table recipe_instructions (
    id SERIAL PRIMARY KEY,
    recipe_id integer,
    step integer NOT NULL,
    instruction text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_recipe
        FOREIGN KEY(recipe_id)
            REFERENCES recipes(id)
);

CREATE table recipe_ingredients (
    id SERIAL PRIMARY KEY,
    recipe_id integer,
    ingredient text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT fk_recipe
        FOREIGN KEY(recipe_id)
            REFERENCES recipes(id)
);

INSERT INTO recipes (name, description) VALUES
('Biryani', 'Make this classic Indian dish for deliciously moist lamb with paneer, rice and spinach, all spiced to perfection. Great for casual entertaining');

INSERT INTO recipe_instructions (id, recipe_id, step, instruction) VALUES
(1, 1, 1, 'Toss the lamb in a bowl with the garlic, ginger and a large pinch of salt. Marinate in the fridge overnight or for at least a couple of hours.'),
(2, 1, 2, 'Heat the oil in a casserole. Fry the lamb for 5-10 mins until starting to brown. Add the onion, cumin seeds and nigella seeds, and cook for 5 mins until starting to soften. Stir in the curry paste, then cook for 1 min more. Scatter in the rice and curry leaves, then pour over the stock and bring to the boil. Meanwhile, heat oven to 180C/160C fan/gas 4.'),
(3, 1, 3, 'Stir in the paneer, spinach and some seasoning. Cover the dish with a tight lid of foil, then put the lid on to ensure it’s well sealed. Cook in the oven for 20 mins, then leave to stand, covered, for 10 mins. Bring the dish to the table, remove the lid and foil, scatter with the coriander and chillies and serve with yogurt on the side.');

INSERT INTO recipe_ingredients (id, recipe_id, ingredient) VALUES
(1, 1, '400g lamb neck, cut into small cubes'),
(2, 1, '4 garlic cloves, grated'),
(3, 1, '1 tbsp finely grated ginger'),
(4, 1, '1 tbsp sunflower oil'),
(5, 1, '1 large onion, chopped'),
(6, 1, '1 tbsp cumin seeds'),
(7, 1, '1 tbsp nigella seeds'),
(8, 1, '1 tbsp Madras spice paste'),
(9, 1, '200g basmati rice, rinsed well'),
(10, 1, '8 curry leaves'),
(11, 1, '400ml good-quality lamb or chicken stock'),
(12, 1, '100g paneer, chopped'),
(13, 1, '200g spinach, cooked and water squeezed out');