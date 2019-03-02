/**
 * The types that can be used in this system.
 * @type {{CHECKBOX: string, NUMBER: string, TEXTAREA: string, RADIO: string, TEXT: string, URL: string}}
 */
const TYPES = {
    CHECKBOX: 'checkbox',
    NUMBER: 'number',
    RADIO: 'radio',
    TEXT: 'text',
    TEXTAREA: 'textarea',
    URL: 'url',
};

/**
 * Regular expression for replaceAll for the prefix for the form
 * @type {RegExp}
 */
const REGEXP = /\W/;

/**
 *
 * @constructor
 */
let Form = function() {
    this.id = 0;

    /**
     * Holds the fields of the form
     * @type {Object[]}
     */
    this.fields = [];

    /**
     * Displaying title of the form
     * @type {string}
     */
    this.title = '';

    /**
     * Holds the name of the form
     * @type {string}
     */
    this.name = '';

    /**
     * Holds the description of the form
     * @type {string}
     */
    this.description = '';

    /**
     * Holds the prefix of the form
     * @type {string}
     */
    this.prefix = '';

    /**
     * Holds the output of the form
     * @type {HTMLElement}
     */
    this.output = null;

    /**
     * Holds the boolean value if the form is weighted
     * @type {boolean}
     */
    this.weighted = false;

    /**
     * Holds the request URL for the form
     * @type {string}
     */
    this.request = '';

    /**
     * Request method
     * @type {string}
     */
    this.method = 'POST';

    /**
     * Initialize some event listeners
     */
    this.Initialize = function(options) {
        this.output = document.querySelector(options.output);
        this.title = options.title;
        this.weighted = options.weighted || false;
        this.request = options.request;
        this.method = options.method || this.method;

        this.renderAll();
    };

    /**
     * Add a new field into the fields array
     * @param {Field} field
     */
    this.add = function(field) {
        this.fields.push(field);
    };

    /**
     * Sorts the field based on the value of 'order'.
     * Lowest first, highest last.
     */
    this.sort = function() {
        this.fields.sort((a, b) => {
            if (a.order < b.order) {
                return -1;
            } else if (a.order > b.order) {
                return 1;
            } else {
                return 0;
            }
        });
    };

    /**
     * Creates a new field in the 'fields' array, then renders all.
     */
    this.NewField = function() {
        this.fields.push(new Field({
            weighted: this.weighted,
        }));

        this.renderAll();
    };

    /**
     * Clean the output HTML, then render all fields in it.
     */
    this.renderAll = function() {
        let container, row, col, formGroup, label, input, textarea, hr, button, accordion, h1, form, hidden, submit;

        this.output.innerHTML = '';

        // <div class="container-fluid">...</div>
        container = document.createElement('div');
        container.classList.add(...['container-fluid']);

        // <div class="row">...</div>
        row = document.createElement('div');
        row.classList.add(...['row']);

        // <div class="col">...</div>
        col = document.createElement('div');
        col.classList.add(...['col']);

        // <h1 class="display-4">...</h1>
        h1 = document.createElement('h1');
        h1.classList.add(...['display-4']);
        h1.innerText = this.title;

        // <hr>
        hr = document.createElement('hr');

        col.appendChild(h1);
        col.appendChild(hr);

        form = document.createElement('form');
        form.setAttribute('method', this.method);
        form.setAttribute('action', this.request);

        col.appendChild(form);

        // <div class="form-group">...</div>
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        // <label for="form_name">Form Name</label>
        label = document.createElement('label');
        label.setAttribute('for', 'form_name');
        label.innerText = 'Form Name';

        // <input type="text" name="form_name" id="form_name" placeholder="Form Name">
        input = document.createElement('input');
        input.setAttribute('type', 'text');
        input.setAttribute('name', 'form_name');
        input.setAttribute('placeholder', 'Form Name');
        input.setAttribute('required', '');
        input.classList.add(...['form-control']);
        input.id = 'form_name';
        input.value = this.name;

        input.addEventListener('keyup', e => {
            this.name = e.target.value;
            this.prefix = this.name.replaceAll(REGEXP, '_').toLowerCase();
            while (this.prefix.includes('__')) {
                this.prefix = this.prefix.replaceAll('__', '_');
            }

            this.fields.forEach((e, i) => {
                e.name = `${this.prefix}_${e.type}_${i}`;
            });
        });

        formGroup.appendChild(label);
        formGroup.appendChild(input);

        form.appendChild(formGroup);

        // <div class="form-group">...</div>
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        // <label for="form_description">Form Description</label>
        label = document.createElement('label');
        label.setAttribute('for', 'form_description');
        label.innerText = 'Form Description';

        // <textarea class="form-control" id="form_description" placeholder="Form Description"></textarea>
        textarea = document.createElement('textarea');
        textarea.classList.add(...['form-control']);
        textarea.setAttribute('placeholder', 'Form Description');
        textarea.id = 'form_description';
        textarea.value = this.description;

        textarea.addEventListener('keyup', e => {
            this.description = e.target.value;
        });

        formGroup.appendChild(label);
        formGroup.appendChild(textarea);

        form.appendChild(formGroup);

        // <hr>
        hr = document.createElement('hr');
        form.appendChild(hr);

        // <button type="button" class="btn btn-dark" id="add_field">Add field</button>
        button = document.createElement('button');
        button.setAttribute('type', 'button');
        button.classList.add(...['btn', 'btn-dark']);
        button.id = 'add_field';
        button.innerText = 'Add field';

        button.addEventListener('click', () => {
            this.NewField();
        });

        form.appendChild(button);

        // <hr>
        hr = document.createElement('hr');
        form.appendChild(hr);

        // <div class="accordion mb-5" id="accordion">...</div>
        accordion = document.createElement('div');
        accordion.classList.add(...['accordion', 'mb-5']);
        accordion.id = 'accordion';

        form.appendChild(accordion);

        // <hr>
        hr = document.createElement('hr');
        form.appendChild(hr);

        // <input type="submit" class="btn btn-primary">
        submit = document.createElement('input');
        submit.setAttribute('type', 'submit');
        submit.classList.add(...['btn', 'btn-primary']);

        submit.addEventListener('click', () => {
            this.prefix = this.name.replaceAll(REGEXP, '_').toLowerCase();
            while (this.prefix.includes('__')) {
                this.prefix = this.prefix.replaceAll('__', '_');
            }

            this.fields.forEach((e, i) => {
                e.name = `${this.prefix}_${e.type}_${i}`;
            });

            hidden.value = this.toJSON();
        });

        form.appendChild(submit);

        // <hr>
        hr = document.createElement('hr');
        form.appendChild(hr);

        this.fields.forEach((element, index) => {
            accordion.appendChild(element.createCard(index));
        });

        // <input type="hidden" name="data">
        hidden = document.createElement('input');
        hidden.setAttribute('type', 'hidden');
        hidden.setAttribute('name', 'data');

        form.appendChild(hidden);

        row.appendChild(col);
        container.appendChild(row);

        this.output.appendChild(container);
    };

    /**
     * Returns the for as an Javascript Object, easy translated to JSON.
     * @returns {object}
     */
    this.toJSON = function() {
        let result = [];

        this.fields.forEach(element => {
            result.push(element.Get());
        });

        return JSON.stringify({
            id:             this.id,
            prefix:         this.prefix,
            name:           this.name,
            description:    this.description,
            fields:         result,
        });
    };

    this.fromJSON = function(data) {
        let json = JSON.parse(data);

        this.id = json.id;
        this.name = json.name;
        this.description = json.description;
        this.prefix = json.prefix;

        for (let i = 0; i < json.fields.length; i++) {
            let f = new Field(this.weighted);
            f.name = json.fields[i].name;
            f.description = json.fields[i].description;
            f.label = json.fields[i].label;
            f.type = json.fields[i].type;
            f.order = json.fields[i].order;
            f.choices = json.fields[i].choices.split(',');

            this.fields.push(f);
        }


        this.renderAll();
    };
};

/**
 * Field is an object that holds data about fields, and functions for displaying it.
 * @param {object} settings
 * @constructor
 */
let Field = function(settings) {
    /**
     * Holds the type of the field
     * @type {string}
     */
    this.type = TYPES.TEXT;

    /**
     * Holds the name of the field
     * @type {string}
     */
    this.name = '';

    /**
     * Holds the label of the field
     * @type {string}
     */
    this.label = '';

    /**
     * Holds the description of the field
     * @type {string}
     */
    this.description = '';

    /**
     * Holds the order of the field
     * @type {number}
     */
    this.order = 0;

    /**
     * Holds the weight of the field
     * @type {number}
     */
    this.weight = 0;

    /**
     * Holds the choices if the type is a multi-choice type
     * @type {Array}
     */
    this.choices = [];

    /**
     * Holds the settings for the field
     * @type {{weighted: (boolean)}}
     */
    this.settings = {
        weighted: settings.weighted || false,
    };

    /**
     * Creates a card with fields, and Bootstrap-styling
     * @param {number} id
     * @returns {HTMLElement}
     */
    this.createCard = function(id) {
        // Declaring all variable names
        let card, cardHeader, cardBody, h2, button, collapse,
            formGroup, label, select, input, option, small, textarea;

        this.name = `${this.type}_${id}`;

        // Defining the card-wrapper
        card = document.createElement('div');
        card.classList.add(...['card']);

        // Defining the card-header
        cardHeader = document.createElement('div');
        cardHeader.classList.add(...['card-header']);
        cardHeader.id = `heading_${id}`;

        // Defining the card-title
        h2 = document.createElement('h2');
        h2.classList.add(...['mb-0']);

        // Defining the accordion button for the card
        button = document.createElement('button');
        button.classList.add(...['btn', 'btn-link']);
        button.setAttribute('type', 'button');
        button.setAttribute('data-toggle', 'collapse');
        button.setAttribute('data-target', `#collapse_${id}`);
        button.setAttribute('aria-controls', `collapse_${id}`);
        button.setAttribute('aria-expanded', 'false');
        button.innerText = `Question ${id+1}`;

        // Appends button into title, and title into card-header
        h2.appendChild(button);
        cardHeader.appendChild(h2);

        // Defining the collapsing wrapper
        collapse = document.createElement('div');
        collapse.classList.add(...['collapse']);
        collapse.id = `collapse_${id}`;
        collapse.setAttribute('aria-labelledby', `header_${id}`);
        collapse.setAttribute('data-parent', '#accordion');

        // Defining the card-body
        cardBody = document.createElement('div');
        cardBody.classList.add(...['card-body']);

        // Appending the card-body into the collapsing wrapper
        collapse.appendChild(cardBody);

        // == INPUT TYPE START ==
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        label = document.createElement('label');
        label.setAttribute('for', `type_${id}`);
        label.innerText = 'Type';

        select = document.createElement('select');
        select.classList.add(...['form-control']);
        select.setAttribute('name', `type_${id}`);
        select.id = `type_${id}`;

        for (let type in TYPES) {
            if (typeof(TYPES[type]) === 'string') {
                option = document.createElement('option');
                option.setAttribute('value', TYPES[type]);
                option.innerText = TYPES[type].toUpperCase();

                if (this.type === TYPES[type]) {
                    option.setAttribute('selected', '');
                } else {
                    option.removeAttribute('selected');
                }

                select.appendChild(option);
            }
        }

        select.addEventListener('change', e => {
            this.type = e.target.value;
            // TODO (Svein): Show choices on change for 'radio'
            if (this.type === TYPES.RADIO) {
                document.getElementById(`choices_form-group_${id}`).classList.remove(...['sr-only']);
            } else {
                document.getElementById(`choices_form-group_${id}`).classList.add(...['sr-only']);
            }

            this.name = `${this.type}_${id}`;
        });

        formGroup.appendChild(label);
        formGroup.appendChild(select);
        cardBody.appendChild(formGroup);
        // == INPUT TYPE END ==

        // == CHOICES START ==
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);
        if (this.type !== TYPES.RADIO) {
            formGroup.classList.add(...['sr-only']);
        }
        formGroup.id = `choices_form-group_${id}`;

        label = document.createElement('label');
        label.setAttribute('for', `choices_${id}`);
        label.innerText = 'Choices';

        textarea = document.createElement('textarea');
        textarea.classList.add(...['form-control']);
        textarea.setAttribute('name', `choices_${id}`);
        textarea.id = `choices_${id}`;
        textarea.value = this.choices.join('\n');

        textarea.addEventListener('keyup', e => {
            this.choices = e.target.value.split('\n');
            let index = this.choices.indexOf('');
            if (index !== -1) {
                this.choices.splice(index, 1);
            }
        });

        small = document.createElement('small');
        small.classList.add(...['form-text', 'text-muted']);
        small.innerText = 'Separate choices per line. Watch out for blank lines!';

        formGroup.appendChild(label);
        formGroup.appendChild(textarea);
        formGroup.appendChild(small);
        cardBody.appendChild(formGroup);
        // == CHOICES END ==

        // == LABEL START ==
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        label = document.createElement('label');
        label.setAttribute('for', `label_${id}`);
        label.innerText = 'Label';

        input = document.createElement('input');
        input.classList.add(...['form-control']);
        input.setAttribute('type', 'text');
        input.setAttribute('placeholder', 'Label');
        input.setAttribute('name', `label_${id}`);
        input.setAttribute('required', '');
        input.id = `label_${id}`;
        input.value = this.label;

        // EventListener to the input-field, saving the value
        input.addEventListener('keyup', e => {
            e.preventDefault();
            this.label = e.target.value;
        });

        formGroup.appendChild(label);
        formGroup.appendChild(input);
        cardBody.appendChild(formGroup);
        // == LABEL END ==

        // == DESCRIPTION START ==
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        label = document.createElement('label');
        label.setAttribute('for', `description_${id}`);
        label.innerText = 'Description';

        textarea = document.createElement('textarea');
        textarea.classList.add(...['form-control']);
        textarea.setAttribute('placeholder', 'Description');
        textarea.setAttribute('placeholder', 'Description');
        textarea.setAttribute('name', `description_${id}`);
        textarea.id = `description_${id}`;
        textarea.value = this.description;

        textarea.addEventListener('keyup', e => {
            this.description = e.target.value;
        });

        formGroup.appendChild(label);
        formGroup.appendChild(textarea);
        cardBody.appendChild(formGroup);
        // == DESCRIPTION END ==

        // == ORDER START ==
        formGroup = document.createElement('div');
        formGroup.classList.add(...['form-group']);

        label = document.createElement('label');
        label.setAttribute('for', `order_${id}`);
        label.innerText = 'Order';

        input = document.createElement('input');
        input.classList.add(...['form-control']);
        input.setAttribute('type', 'number');
        input.setAttribute('min', '0');
        input.setAttribute('name', `order_${id}`);
        input.setAttribute('required', '');
        input.id = `order_${id}`;
        input.value = this.order;
        // EventListener to the input-field, saving the value
        input.addEventListener('keyup', e => {
            e.preventDefault();
            this.order = parseInt(e.target.value);
        });

        small = document.createElement('small');
        small.classList.add(...['form-text', 'text-muted']);
        small.innerText = 'Ordered ascending by number. Smallest first. Minimum value: 0.';

        formGroup.appendChild(label);
        formGroup.appendChild(input);
        formGroup.appendChild(small);
        cardBody.appendChild(formGroup);
        // == ORDER END ==

        // Check if form is weighted
        if (this.settings.weighted) {
            // == WEIGHT START ==
            formGroup = document.createElement('div');
            formGroup.classList.add(...['form-group']);

            label = document.createElement('label');
            label.setAttribute('for', `weight_${id}`);
            label.innerText = 'Weight';

            input = document.createElement('input');
            input.classList.add(...['form-control']);
            input.setAttribute('type', 'number');
            input.setAttribute('min', '0');
            input.setAttribute('name', `weight_${id}`);
            input.id = `weight_${id}`;
            input.value = this.weight;
            // EventListener to the input-field, saving the value
            input.addEventListener('keyup', e => {
                e.preventDefault();
                this.weight = parseInt(e.target.value);
            });

            formGroup.appendChild(label);
            formGroup.appendChild(input);
            cardBody.appendChild(formGroup);
            // == WEIGHT END ==
        }

        card.appendChild(cardHeader);
        card.appendChild(collapse);

        return card;
    };

    /**
     * Returns the Field-object as a simple object
     * @returns {object}
     */
    this.Get = function() {
        return {
            type:           this.type,
            name:           this.name,
            label:          this.label,
            description:    this.description,
            order:          this.order,
            weight:         this.weight,
            choices:        this.choices.join(','),
        }
    };
};

/**
 * A replace-function that takes all replacements, not just the first.
 * @param {string|RegExp} search
 * @param {string} replacement
 * @returns {string}
 */
String.prototype.replaceAll = function(search, replacement) {
    let target = this;
    return target.replace(new RegExp(search, 'g'), replacement);
};
