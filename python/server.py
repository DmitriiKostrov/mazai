from bottle import post, route, run, template, request, static_file, error, redirect
from database_json import init_trips, get_all_trips, join_trip, left_trip, add_trip, remove_trip, edit_trip

@route('/')
def index():
    trips = get_all_trips()
    #return template('<b>Hello {{name}}</b>!', name=name)
    #return trips#get_tomorrow_trips()
    return template('templates/index.html', trips=trips)
    #, name=request.environ.get('REMOTE_ADDR'))

@route('/<day>')
def index():
    trips = get_all_trips()
    #return template('<b>Hello {{name}}</b>!', name=name)
    #return trips#get_tomorrow_trips()
    return template('templates/day.html', trips=trips, day=int(day))

@route('/trip/<id>/<day>/join')
def join(id, day):
    join_trip(int(id), int(day))
    redirect('/')

@route('/trip/<id>/<day>/left')
def join(id, day):
    left_trip(int(id), int(day))
    redirect('/')

@route('/trip/<id>/remove')
def remove(id):
    remove_trip(int(id))
    redirect('/')

@route('/trip/new')
def new():
    #return template('<b>Hello {{name}}</b>!', name=name)
    #return trips#get_tomorrow_trips()
    trip = {'time': '18:10', 'delay': '10', 
            'driver': '',
            'to': '', 
            'from': '', 
            'desc': '', 
            'days' : [ 
                {'seats': 4, 'guests': 0},
                {'seats': 4, 'guests': 0},
                {'seats': 4, 'guests': 0},
                {'seats': 4, 'guests': 0},
                {'seats': 4, 'guests': 0},
                {'seats': 0, 'guests': 0},
                {'seats': 0, 'guests': 0}
            ]
        }
    return template('templates/new.html', trip=trip, id=-1)

@post('/trip/edit')
def edit():
    time   = request.forms.time
    delay  = request.forms.delay
    driver = request.forms.driver
    desc   = request.forms.desc
    to     = request.forms.to
    frm    = request.forms.frm # from
    seats0 = request.forms.seats0
    seats1 = request.forms.seats1
    seats2 = request.forms.seats2
    seats3 = request.forms.seats3
    seats4 = request.forms.seats4
    seats5 = request.forms.seats5
    seats6 = request.forms.seats6
    id     = request.forms.id

    if frm and to and driver and time:
        trip = {'driver': driver, 'time': time, 'desc': desc, 
            'to': to, 
            'from': frm, 
            'delay': delay,
            'days' : [ 
                {'seats': int(seats0), 'guests': 0},
                {'seats': int(seats1), 'guests': 0},
                {'seats': int(seats2), 'guests': 0},
                {'seats': int(seats3), 'guests': 0},
                {'seats': int(seats4), 'guests': 0},
                {'seats': int(seats5), 'guests': 0},
                {'seats': int(seats6), 'guests': 0}
            ]
        }
        if id == '-1':
            add_trip(trip)
        else:
            edit_trip(trip, int(id))
    
    redirect('/')

@route('/trip/<id>/modify')
def modify(id):
    trips = get_all_trips()
    return template('templates/new.html', trip=trips[int(id)], id=id)

@error(404)
def error404(error):
    return 'Nothing here, sorry'

@error(505)
def error505(error):
    return 'Nothing here, sorry, hehe'

init_trips()
run(host='0.0.0.0', port=8585, debug=True)
