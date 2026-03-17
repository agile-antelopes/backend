from fastapi import FastAPI
import pyscopg2
from pydantic import BaseModel

app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: str | None = None):
    return {"item_id": item_id, "q": q}


app = FastAPI()


class Item(BaseModel):
    name: str
    price: float
    is_offer: bool | None = None


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: str | None = None):
    return {"item_id": item_id, "q": q}


@app.put("/items/{item_id}")
def update_item(item_id: int, item: Item):
    return {"item_name": item.name, "item_id": item_id}





class Item(BaseModel):
    name: str
    description: str | None = None

app = FastAPI()
router = APIRouter()

@router.post("/items/")
def create_item(item: Item):
    return {"message": "Item created"}

app.include_router(router)



class Countries(BaseModel):
    # Info
    country_name: str
    languages: list[str]
    currency = str
    government_type = str

    # Background and geography
    location = str
    geographic_features = list[str]
    climate = str
    history = str
    major_cities = list[str]
    transportation = str

    # Values and Beliefs
    religion = list[str]
    values = list[str]
    customs_and_traditions = str
    common_misunderstandings = str

    # Social Life
    food = list[str]
    greetings_and_social_etiquette = str
    common_phrases_and_slang = list[str]

    # Safety
    safety_tips = list[str]
    laws = list[str]
    emergency_services = list[str]

    # Social Do's and Don'ts
    social_dos = list[str]
    social_donts = list[str]

class Summaries(BaseModel):
    copywriter = str
    date = str
    country_id = str
    continent = str

class Interviews(BaseModel):
    date = str
    interviewer = str
    interviewee = str
    country_id = str
    id = str

class A_Q_and_A_Background_and_Location(BaseModel):
    # Get from form create class for each section
    where_are_you_from = str
    major_historcal_elements = list[str]

class B_Q_and_A_Core_Cultural_Values_and_Beliefs(BaseModel):
    key_values = list[str]
    religious_views_and_practices = list[str]
    cultural_misunderstandings = list[str]

class C_Q_and_A_Daily_Life_and_Society(BaseModel):
    normal_day = str
    expectations_areound_meals = list[str]
    transportation_public = str
    transportation_driving = str
    transportation_walking = str
    geographic_and_climate_factors_and_how_to_dress_appropriately = str
    traditional_clothing = str

class D_Q_and_A_Greetings_Manners_Social_Etiquette(BaseModel):
    standard_greeting_strangers = str
    standard_greeting_friends = str
    good_manners_in_public_and_meeting_new_people = str
    social_dos = list[str]
    social_donts = list[str]
    topical_taboos = list[str]
    behavioral_taboos = list[str]

class E_Q_and_A_Food_and_Traditions(BaseModel):
    common_dishes = list[str]
    unique_traditions = list[str]

class F_Q_and_A_Practical_Travel_Considerations(BaseModel):
    laws = list[str]
    safety_concerns_and_dangers = list[str]

class G_Q_and_A_Language_and_Slang(BaseModel):
    slang_words = list[str]
    slang_word_definitions = list[str]
    languages = list[str]
    common_phrases = list[str]

class H_Q_and_A_Cultural_Adjustment(BaseModel):
    cultutal_differences = list[str]
    culture_shock = str
    hardest_cultural_adjustment = str
    cultural_practices = str

class I_Q_and_A_Reflection_and_Advice(BaseModel):
    advice = str
    know_ahead_of_time = str
    one_thing_to_know = str

class J_Q_and_A_Follow_Up(BaseModel):
    contacted_by_users = str
    know_anyone_else = str


@app.post("/create/country")
def create_item(country: Countries):
    
    print(country)

@app.post("/create/summaries")
def create_item(summary: Summaries):
    
    print(summary)

@app.post("/create/interviews")
def create_item(interview: Interviews):
    
    print(interview)

@app.post("/create/background_and_locations")
def create_item(background_and_location: A_Q_and_A_Background_and_Location):
    
    print(background_and_location)

@app.post("/create/core_cultural_values_and_beliefs")
def create_item(core_cultural_values_and_beliefs: B_Q_and_A_Core_Cultural_Values_and_Beliefs):
    
    print(core_cultural_values_and_beliefs)

@app.post("/create/daily_life_and_society")
def create_item(daily_life_and_society: C_Q_and_A_Daily_Life_and_Society):
    
    print(daily_life_and_society)

@app.post("/create/greetings_manners_social_etiquette")
def create_item(greetings_manners_social_etiquette: D_Q_and_A_Greetings_Manners_Social_Etiquette):
    
    print(greetings_manners_social_etiquette)

@app.post("/create/food_and_traditions")
def create_item(food_and_traditions: E_Q_and_A_Food_and_Traditions):
    
    print(food_and_traditions)

@app.post("/create/practical_travel_considerationsaditions")
def create_item(practical_travel_considerationsaditions: F_Q_and_A_Practical_Travel_Considerations):
    
    print(practical_travel_considerationsaditions)

@app.post("/create/cultural_adjustment")
def create_item(language_and_slang: H_Q_and_A_Cultural_Adjustment):
    
    print(language_and_slang)