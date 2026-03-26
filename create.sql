CREATE TABLE country (
                         country_code CHAR(3) PRIMARY KEY,
                         country_name VARCHAR NOT NULL,
                         facts TEXT
);

CREATE TABLE "topic_tag-A" (
                               topic_tag_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               topic_tag VARCHAR NOT NULL
);

-- (Tabla antigua o alternativa que aparecía en tu imagen)
CREATE TABLE article (
                         article_id VARCHAR PRIMARY KEY,
                         content TEXT
);

-- ==========================================
-- TABLAS DEPENDIENTES (Con Foreign Keys)
-- ==========================================

CREATE TABLE "interview-A" (
                               interview_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                               date DATE NOT NULL,
                               interviewer_name VARCHAR NOT NULL,
                               interviewee_name VARCHAR NOT NULL,
                               country_id CHAR(3) REFERENCES country(country_code)
);

CREATE TABLE "response-A" (
                              response_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                              question TEXT NOT NULL,
                              answer TEXT,
                              interview_id INTEGER REFERENCES "interview-A"(interview_id) ON DELETE CASCADE,
                              country_id VARCHAR(3) REFERENCES country(country_code),
                              topic_tag_id INTEGER REFERENCES "topic_tag-A"(topic_tag_id)
);

CREATE TABLE "article-A" (
                             article_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                             copywriter_name VARCHAR NOT NULL,
                             date DATE NOT NULL,
                             country_id VARCHAR(3) REFERENCES country(country_code),
                             content TEXT,
                             topic_tag_id INTEGER REFERENCES "topic_tag-A"(topic_tag_id)
);

CREATE TABLE "quote-A" (
                           quote_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                           article_id INTEGER REFERENCES "article-A"(article_id) ON DELETE CASCADE,
                           response_id INTEGER REFERENCES "response-A"(response_id) ON DELETE CASCADE
);

CREATE TABLE worldloom.question_template (
                                             question_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                                             topic_tag_id INTEGER REFERENCES worldloom."topic_tag-A"(topic_tag_id) ON DELETE CASCADE,
                                             question_text TEXT NOT NULL
);

-- Asegúrate de estar en el esquema correcto
SET search_path TO worldloom;

INSERT INTO question_template (topic_tag_id, question_text) VALUES
-- A. Background & Location (ID: 1)
(1, 'Where are you from (or where did you serve your mission) specifically? (city/region and country.'),
(1, 'What are some major historical elements of your country or culture?'),

-- B. Core Cultural Values & Beliefs (ID: 2)
(2, 'What are the top key values in your culture?'),
(2, 'What are the most common religious views or practices in your culture?'),
(2, 'What are some commonly misunderstood things about your culture?'),

-- C. Daily Life & Society (ID: 3)
(3, 'What does a normal day look like in your culture?'),
(3, 'What are expectations around family gatherings and sharing meals?'),
(3, 'How does transportation work (public transit, driving, walking)?'),
(3, 'What are important geographic or climate factors that shape daily life, and how should you dress for the weather?'),
(3, 'What does traditional clothing look like, and when is it worn today?'),

-- D. Greetings, Manners & Social Etiquette (ID: 4)
(4, 'What does a standard greeting look like (strangers vs. close friends)?'),
(4, 'What is considered good manners in public or when meeting new people?'),
(4, 'What are important do’s and don’ts in your culture?'),
(4, 'Are there topics or behaviors that are considered taboo in conversation?'),

-- E. Food & Traditions (ID: 5)
(5, 'What are common dishes?'),
(5, 'What unique traditions might surprise someone from the outside?'),

-- F. Practical Travel Considerations (ID: 6)
(6, 'Are there any unique laws visitors should be aware of?'),
(6, 'Are there any safety concerns or dangers foreigners should understand?'),

-- G. Language & Slang (ID: 7)
(7, 'What are some common slang words or informal expressions, and what do they mean?'),
(7, 'What language(s) are spoken there? What are a few common phrases visitors should know?'),

-- H. Cultural Adjustment Experience (BYUI / Mission Context) (ID: 8)
(8, 'What cultural differences did you notice when coming to BYUI or your mission area?'),
(8, 'What shocked you most culturally?'),
(8, 'What was the hardest cultural adjustment?'),
(8, 'What is something normal in your home culture that people don’t typically do here?'),

-- I. Reflection & Advice (ID: 9)
(9, 'What advice would you give someone experiencing your culture for the first time?'),
(9, 'If you were preparing to visit a completely new culture, what would you want to know ahead of time?'),
(9, 'If someone visited your country today, what is one thing they should know that we haven''t talked about?'),

-- J. Follow-Up (ID: 10)
(10, 'Would you be willing to be contacted by users of this website? If yes, please provide contact information.'),
(10, 'Do you know anyone else who would be willing to answer these questions?');