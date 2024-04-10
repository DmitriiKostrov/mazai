import os, json
from datetime import datetime, timedelta

MAIN_FILE=os.environ.get('dataFile', 'trips.json')
TODAY_CACHE=[]

def init_trips():
    get_today_trips()

def get_trips_from_file(file_name):
    try:
        data = []
        with open(file_name) as f:
            data = json.loads(f.read())
    except e:
        print('%s' % e)
    finally:
        return data

def get_trips(when):
    return get_trips_from_file(when.strftime('%Y%m%d') + '.json')

def get_today_trips():
    global TODAY_CACHE
    if TODAY_CACHE == []:
        now = datetime.now()
        TODAY_CACHE=get_trips(now)

    return TODAY_CACHE

def get_all_trips():
    global TODAY_CACHE
    if TODAY_CACHE == []:
        TODAY_CACHE=get_trips_from_file(MAIN_FILE)

    return TODAY_CACHE

def get_tomorrow_trips():
    tomorrow = datetime.now() + timedelta(days=1)
    return get_trips(tomorrow)

def update_trips(trips):
    with open(MAIN_FILE, 'w') as f:
        json.dump(trips, f, indent=4, sort_keys=False, ensure_ascii=False)

def add_trip(trip):
    TODAY_CACHE.append(trip)
    update_trips(TODAY_CACHE)

def join_trip(trip_id, day):
    if trip_id >= len(TODAY_CACHE):
        raise Exception('Invalid Trip ID')

    trip = TODAY_CACHE[trip_id]
    if len(trip['days']) <= day:
        raise Exception('Invalid day number: %d' % day) 
    
    if trip['days'][day]['guests'] >= trip['days'][day]['seats']:
        raise Exception('No more seats for this Trip')
    trip['days'][day]['guests'] += 1

    TODAY_CACHE[trip_id] = trip
    update_trips(TODAY_CACHE)

def left_trip(trip_id, day):
    if trip_id >= len(TODAY_CACHE):
        raise Exception('Invalid Trip ID')

    trip = TODAY_CACHE[trip_id]
    if len(trip['days']) <= day:
        raise Exception('Invalid day number: %d' % day) 
    
    if trip['days'][day]['guests'] <= 0:
        raise Exception('No more seats for this Trip')
    trip['days'][day]['guests'] -= 1

    TODAY_CACHE[trip_id] = trip
    update_trips(TODAY_CACHE)

def remove_trip(trip_id):
    if trip_id >= len(TODAY_CACHE):
        raise Exception('Invalid Trip ID')

    del TODAY_CACHE[trip_id]
    update_trips(TODAY_CACHE)

def edit_trip(trip, trip_id):
    if trip_id >= len(TODAY_CACHE):
        raise Exception('Invalid Trip ID')
    
    # fix new guests number
    prev_trip = TODAY_CACHE[trip_id]
    for idx, day in enumerate(prev_trip['days']):
        old_guests = int(day['guests'])
        new_seats = trip['days'][idx]['seats']
        if old_guests > new_seats:
            old_guests = new_seats
        trip['days'][idx]['guests'] = old_guests

    TODAY_CACHE[trip_id] = trip
    update_trips(TODAY_CACHE)
