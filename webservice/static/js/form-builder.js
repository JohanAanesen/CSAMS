/**
 * A simple application for creating dynamic forms.
 * Using Bootstrap for styling.
 * Supports these field types:
 *      - text
 *      - textarea
 *      - number
 *      - url
 *      - checkbox
 *      - radio
 *
 * How to use:
 *
 * DEPENDENCIES:
 *      - Bootstrap CSS, tested only on 4.2.1, should work with all 4.x
 *      - https://cdn.jsdelivr.net/npm/sortablejs@latest/Sortable.min.js
 *
 * HTML
 <div id="output"></div>

 * JS:
 let formBuilder = new FormBuilder({
    title: 'Form Builder', // Title of the form
    description: 'This is some description', // Description of the form
    request: '/form-builder', // Action on the form
    method: 'POST', // Method of the form
    output: document.getElementById('output'), // The output element of the application
    weighted: true, // Boolean: is the form weighted
});


 * @author Svein Are Danielsen
 * @version 1.0.0
 * @licence GPL-3.0
 *
 * TODO:
 *  - Add button to remove elements
 */

/**
 * FormBuilder holds the form, and initialize it
 * @param {object} args
 * @constructor
 */
function FormBuilder(args) {
    this.form = new Form(args);
    this.form.initialize();

    /**
     * Imports data from a string in JSON-format
     * @param {string} data
     */
    this.import = function(data) {
        this.form.fromJSON(data);
    };
}

/**
 * Input field types
 * @type {{TEXT: string, TEXTAREA: string, URL: string, NUMBER: string, CHECKBOX: string, RADIO: string}}
 */
const TYPES = {
    TEXT: 'text',
    TEXTAREA: 'textarea',
    URL: 'url',
    NUMBER: 'number',
    CHECKBOX: 'checkbox',
    RADIO: 'radio',
};

/**
 * Form object holds all form data and all fields
 * @param {object} args
 * @constructor
 */
function Form(args) {
    this.id = (args.id !== undefined) ? args.id : 0;
    this.title = (args.title !== undefined) ? args.title : '';
    this.description = (args.description !== undefined) ? args.description : '';
    this.name = (args.name !== undefined) ? args.name : '';
    this.prefix = (args.prefix !== undefined) ? args.prefix : '';
    this.weighted = (args.weighted !== undefined) ? args.weighted : false;
    this.output = (args.output !== undefined) ? args.output : null;
    this.request = (args.request !== undefined) ? args.request : '';
    this.method = (args.method !== undefined) ? args.method : 'POST';
    this.fields = (args.fields !== undefined) ? args.fields : [];
    this.regexp = (args.regexp !== undefined) ? args.regexp : /\W/;
    this.deleteRequest = (args.deleteRequest !== undefined) ? args.deleteRequest : '';
    this.deleteMethod = (args.deleteMethod !== undefined) ? args.deleteMethod : 'DELETE';

    this.totalWeight = 0;

    this.sortable = null;

    this.lastId = 0;

    /**
     * Initialize the form object
     */
    this.initialize = function() {
        this.render();
    };

    /**
     * Renders the whole system
     */
    this.render = function() {
        this.clearOutput();

        let container = createElement({
            classList: ['container-fluid'],
        });

        let form = createElement({
            type: 'form',
            attributes: [
                {
                    name: 'action',
                    value: this.request,
                },
                {
                    name: 'method',
                    value: this.method,
                },
            ],
        });

        container.appendChild(form);

        form.appendChild(this.renderHeader());

        let hidden = createElement({
            type: 'input',
            attributes: [
                {
                    name: 'type',
                    value: 'hidden',
                },
                {
                    name: 'name',
                    value: 'form_data',
                }
            ],
            id: 'form_data',
        });

        form.appendChild(hidden);

        let row = createElement({
            classList: ['row'],
        });

        row.appendChild(this.renderLeftColumn());
        row.appendChild(this.renderRightColumn());

        form.appendChild(row);

        this.output.appendChild(container);

        this.postRender();
    };

    /**
     * Renders the header of the program, displaying simple information about the form
     * @return {Element}
     */
    this.renderHeader = function() {
        let row = createElement({
            classList: ['row'],
        });

        let col = createElement({
            classList: ['col'],
        });

        let title = createElement({
            type: 'h1',
            classList: ['display-4'],
            innerHTML: this.title + ' <small id="form_name"></small>',
        });

        let description = createElement({
            type: 'p',
            classList: ['lead'],
            innerText: this.description,
        });

        let hr = createElement({type:'hr'});

        col.appendChild(title);
        col.appendChild(description);
        col.appendChild(hr);

        row.appendChild(col);

        return row;
    };

    /**
     * Renders the left side of the software, containing the "new field" button, submit form button,
     * and input field for naming the form
     * @return {Element}
     */
    this.renderLeftColumn = function() {
        let leftCol = createElement({
            classList: ['col-md-3', 'mb-3'],
        });

        let formGroup = createElement({
            classList: ['form-group'],
        });

        let nameLabel = createElement({
            type: 'label',
            attributes: [
                {
                    name: 'for',
                    value: 'form_name',
                }
            ],
            innerText: 'Form Name',
        });

        let nameInput = createElement({
            type: 'input',
            classList: ['form-control'],
            id: 'form_name',
            attributes: [
                {
                    name: 'placeholder',
                    value: 'Form Name',
                },
                {
                    name: 'required',
                    value: '',
                }
            ],
            value: this.name.trim(),
        });

        nameInput.addEventListener('keyup', () => {
            this.name = nameInput.value.trim();
            this.fixPrefix();

            this.postRender();
        });

        formGroup.appendChild(nameLabel);
        formGroup.appendChild(nameInput);

        leftCol.appendChild(formGroup);

        let hr = createElement({ type: 'hr' });

        leftCol.appendChild(hr);

        let newButton = createElement({
            type: 'button',
            classList: ['btn', 'btn-primary', 'btn-block', 'btn-lg'],
            attributes: [
                {
                    name: 'type',
                    value: 'button',
                },
            ],
            innerText: 'Add a new Field',
        });

        newButton.addEventListener('click', () => {
            this.newField();
        });

        hr = createElement({ type: 'hr' });

        leftCol.appendChild(newButton);
        leftCol.appendChild(hr);

        let submitButton = createElement({
            type: 'button',
            classList: ['btn', 'btn-outline-success', 'btn-block'],
            attributes: [
                {
                    name: 'type',
                    value: 'submit',
                },
            ],
            innerText: 'Submit',
            id: 'submit_btn',
        });

        hr = createElement({ type: 'hr' });

        leftCol.appendChild(submitButton);
        leftCol.appendChild(hr);

        if (this.deleteRequest !== '') {
            let deleteButton = createElement({
                type: 'button',
                classList: ['btn', 'btn-outline-danger', 'btn-block'],
                attributes: [
                    {
                        name: 'type',
                        value: 'button',
                    },
                ],
                innerText: 'Delete',
                id: 'delete_btn',
            });

            deleteButton.addEventListener('click', e => {
                let ok = window.confirm('Are you sure you want to delete this?');
                if (ok) {
                    fetch(this.deleteRequest, {
                        method: this.deleteMethod,
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({
                            id: this.id,
                        }),
                    })
                        .then((response) => {
                            if (response.status === 200) {
                                return response.json();
                            } else {
                                return {};
                            }
                        })
                        .then((json) => {
                            if (json.location !== undefined) {
                                window.location = json.location;
                            }
                        })
                        .catch((error) => {
                            console.error(error);
                        })
                }
            });

            hr = createElement({ type: 'hr' });

            leftCol.appendChild(deleteButton);
            leftCol.appendChild(hr);
        }

        return leftCol;
    };

    /**
     * Renders the "right" column of the program, containing the sortable (drag and drop) list of fields.
     * @return {Element}
     */
    this.renderRightColumn = function() {
        let rightCol = createElement({
            classList: ['col-md-9'],
            id: 'sort_me',
        });

        this.fields.forEach((element, index) => {
            rightCol.appendChild(element.render(index));
        });

        //noinspection JSUnresolvedVariable
        if (Sortable !== null) {
            let fields = this.fields;

            //noinspection JSUnresolvedVariable,JSUnresolvedFunction
            Sortable.create(rightCol, {
                handle: '.card-header',
                onEnd: function(e) {
                    let oldIndex = e.oldIndex;
                    let newIndex = e.newIndex;

                    let temp = fields[oldIndex].order;
                    fields[oldIndex].order = fields[newIndex].order;
                    fields[newIndex].order = temp;
                },
            });

            this.fields = fields;
            //this.sortFields(); // TODO (Svein): This cause some bug
        }

        return rightCol;
    };

    /**
     * Handles all rendering that needs to happen at the end of rendering
     */
    this.postRender = function() {
        let formName = document.getElementById('form_name');

        if (/\S/.test(this.name)) {
            formName.innerText = ' - ' + this.name;
        } else {
            formName.innerText = '';
        }

        let submitButton = document.getElementById('submit_btn');

        submitButton.addEventListener('click', () => {
            let hidden = document.getElementById('form_data');
            hidden.value = this.toJSON();
        });


        let self = this;
        this.fields.forEach(e => {
            let deleteButton = document.getElementById(`delete_field_${e.id}`);
            deleteButton.addEventListener('click', () => {
                self.fields = self.fields.filter((f) => {
                    return f.id !== e.id;
                });
                this.render();
            });
        });
    };

    /**
     * Adds up all weight to a total weight, and stores it in the localStorage.
     */
    this.fixTotalWeight = function() {
        this.totalWeight = 0;
        this.fields.forEach(e => {
            this.totalWeight += parseInt(e.weight);
        });
        localStorage.setItem('totalWeight', this.totalWeight);
    };

    /**
     * Clears the output
     */
    this.clearOutput = function() {
        this.output.innerHTML = '';
    };

    /**
     * Creates a new field, and renders the screen
     */
    this.newField = function() {
        this.fields.push(new Field({
            id: this.lastId++,
            order: this.fields.length,
            weighted: this.weighted,
        }));

        this.render();
    };

    /**
     * fixPrefix creates the prefix for the form, replaces all non alphanumeric characters with '_' (underscore)
     * to lower case. Also checks for double '_' (underscores) and replace them with single '_' (underscore).
     */
    this.fixPrefix = function() {
        this.prefix = this.name.trim().replaceAll(this.regexp, '_').toLowerCase();
        while (this.prefix.includes('__')) {
            this.prefix = this.prefix.replaceAll('__', '_');
        }
    };

    /**
     * Returns the prefix for the form
     * @return {string}
     */
    this.getPrefix = function() {
        this.fixPrefix();
        return this.prefix;
    };

    /**
     * Returns the form as a JSON-string
     * @return {string}
     */
    this.toJSON = function() {
        let fields = [];

        for (let i = 0; i < this.fields.length; i++) {
            let f = this.fields[i];
            f.name = `${this.getPrefix()}_${f.type}_${f.id}`;
            fields.push(f.get());
        }

        return JSON.stringify({
            id:         this.id,
            prefix:     this.getPrefix(),
            name:       this.name,
            weighted:   this.weighted,
            fields:     fields,
        });
    };

    /**
     * Convert JSON to a form
     * @param {string} data
     */
    this.fromJSON = function(data) {
        let json = JSON.parse(data);

        this.id = json.id;
        this.name = json.name;
        this.prefix = json.prefix;
        this.weighted = false;

        for (let i = 0; i < json.fields.length; i++) {
            let field = json.fields[i];

            if (field.weight !== 0) {
                this.weighted = true;
            }

            let temp = new Field({
                id: field.id,
                label: field.label,
                name: field.name,
                description: field.description,
                hasComment: field.hasComment,
                type: field.type,
                weight: field.weight,
                order: field.order,
            });

            if (json.weighted) {
                temp.weight = json.fields[i].weight;
            }

            this.fields.push(temp);
        }

        if (this.weighted) {
            this.fields.forEach(e => {
                e.weighted = true;
            })
        }

        this.sortFields();
        this.render();
    };

    /**
     * Sorts fields based on their order-value
     */
    this.sortFields = function() {
        this.fields.sort((a, b) => {
            if (a.order > b.order) {
                return 1;
            } else if (a.order < b.order) {
                return -1;
            } else {
                return 0;
            }
        });
    };
}

/**
 * Field type holds all data for a single field
 * @param {object} args
 * @constructor
 */
function Field(args) {
    this.id = (args.id !== undefined) ? args.id : -1;
    this.type = (args.type !== undefined) ? args.type : TYPES.TEXT;
    this.name = (args.name !== undefined) ? args.name : '';
    this.description = (args.description !== undefined) ? args.description : '';
    this.label = (args.label !== undefined) ? args.label : '';
    this.hasComment = (args.hasComment !== undefined) ? args.hasComment : false;
    this.order = (args.order !== undefined) ? args.order : 0;
    this.weighted = (args.weighted !== undefined) ? args.weighted : false;
    this.weight = (args.weight !== undefined) ? args.weight : 0;
    this.choices = (args.choices !== undefined) ? args.choices : '';
    this.expanded = false;
    this.displayChoices = false;

    /**
     * Render the card
     * @return {Element}
     */
    this.render = function() {
        let card = createElement({
            classList: ['card', 'mb-3'],
            attributes: [
                {
                    name: 'data-order',
                    value: this.order,
                },
                {
                    name: 'data-id',
                    value: this.id,
                },
            ],
        });

        card.appendChild(this.renderHeader());
        card.appendChild(this.renderBody());

        return card;
    };

    /**
     * Render the card header
     * @return {Element}
     */
    this.renderHeader = function() {
        let cardHeader = createElement({
            classList: ['card-header'],
            id: `header_${this.id}`,
        });

        let cardTitle = createElement({
            type: 'h5',
            classList: ['mb-0'],
        });

        let cardButtonClasses = ['btn', 'btn-link'];
        if (!this.expanded) {
            cardButtonClasses.push('collapsed');
        }

        let cardButton = createElement({
            type: 'button',
            classList: cardButtonClasses,
            attributes: [
                {
                    name: 'type',
                    value: 'button',
                },
                {
                    name: 'data-toggle',
                    value: 'collapse',
                },
                {
                    name: 'data-target',
                    value: `#collapse_${this.id}`,
                },
                {
                    name: 'aria-expanded',
                    value: this.expanded,
                },
                {
                    name: 'aria-controls',
                    value: `collapse_${this.id}`,
                },
            ],
            innerHTML: (/\S/.test(this.label) ? this.label : `Question ${this.id}`),
            id: `header_btn_${this.id}`,
        });

        cardButton.addEventListener('click', () => {
            let temp = document.getElementById(`collapse_${this.id}`);
            setTimeout(() => {
                this.expanded = temp.classList.contains('show');
            }, 375);
        });

        cardTitle.appendChild(cardButton);
        cardHeader.appendChild(cardTitle);

        return cardHeader;
    };

    /**
     * Render the card body
     * @return {Element}
     */
    this.renderBody = function() {
        let collapseClasses = ['collapse'];

        if (this.expanded) {
            collapseClasses.push('show');
        }

        let collapse = createElement({
            classList: collapseClasses,
            id: `collapse_${this.id}`,
            attributes: [
                {
                    name: 'aria-labelledby',
                    value: `collapse_${this.id}`,
                },
            ],
        });

        let cardBody = createElement({
            classList: ['card-body'],
        });


        cardBody.appendChild(this.renderTypeInput());
        cardBody.appendChild(this.renderLabelInput());
        cardBody.appendChild(this.renderDescriptionInput());
        cardBody.appendChild(this.renderCommentCheckbox());
        cardBody.appendChild(this.renderChoicesTextarea());

        if (this.weighted) {
            cardBody.appendChild(this.renderWeightInput());
        }

        let deleteButton = createElement({
            type: 'button',
            classList: ['btn', 'btn-danger', 'btn-sm', 'px-2'],
            attributes: [
                {
                    name: 'type',
                    value: 'button',
                },
            ],
            id: `delete_field_${this.id}`,
            innerText: 'Remove this field',
        });

        cardBody.appendChild(deleteButton);

        collapse.appendChild(cardBody);

        return collapse;
    };

    /**
     * Render the select type input
     * @return {Element}
     */
    this.renderTypeInput = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
        });

        let label = createElement({
            type: 'label',
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Type',
            attributes: [
                {
                    name: 'for',
                    value: `type_${this.id}`,
                }
            ]
        });

        formGroup.appendChild(label);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let select = createElement({
            type: 'select',
            classList: ['form-control'],
            id: `type_${this.id}`,
        });

        for (let type in TYPES) {
            let attributes = [
                {
                    name: 'value',
                    value: TYPES[type],
                },
            ];

            if (TYPES[type] === this.type) {
                attributes.push({
                    name: 'selected',
                    value: '',
                });
            }

            let option = createElement({
                type: 'option',
                attributes: attributes,
                innerText: TYPES[type].toUpperCase(),
            });

            select.appendChild(option);
        }

        select.addEventListener('change', e => {
            this.type = e.target.value;
            this.displayChoices = this.type === TYPES.RADIO;

            if (this.displayChoices) {
                document.getElementById(`choices_display_${this.id}`).classList.remove('sr-only');
            } else {
                document.getElementById(`choices_display_${this.id}`).classList.add('sr-only');
            }
        });

        rightSide.appendChild(select);
        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Render thee input
     * @return {Element}
     */
    this.renderLabelInput = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
        });

        let label = createElement({
            type: 'label',
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Label',
            attributes: [
                {
                    name: 'for',
                    value: `label_${this.id}`,
                }
            ]
        });

        formGroup.appendChild(label);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let input = createElement({
            type: 'input',
            classList: ['form-control'],
            attributes: [
                {
                    name: 'type',
                    value: 'text',
                },
                {
                    name: 'placeholder',
                    value: 'Label',
                },
                {
                    name: 'required',
                    value: '',
                }
            ],
            id: `label_${this.id}`,
            value: this.label,
        });

        input.addEventListener('keyup', () => {
            this.label = input.value;
            let temp = document.getElementById(`header_btn_${this.id}`);

            if (!/\S/.test(input.value)) {
                temp.innerText = `Question ${this.id}`;
            } else {
                temp.innerText = input.value;
            }
        });

        rightSide.appendChild(input);
        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Render the description input
     * @return {Element}
     */
    this.renderDescriptionInput = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
        });

        let label = createElement({
            type: 'label',
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Description',
            attributes: [
                {
                    name: 'for',
                    value: `description_${this.id}`,
                }
            ]
        });

        formGroup.appendChild(label);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let input = createElement({
            type: 'textarea',
            classList: ['form-control'],
            attributes: [
                {
                    name: 'placeholder',
                    value: 'Description',
                }
            ],
            id: `description_${this.id}`,
            value: this.description,
        });

        input.addEventListener('keyup', () => {
            this.description = input.value;
        });

        rightSide.appendChild(input);
        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Render the enable comment checkbox input
     * @return {Element}
     */
    this.renderCommentCheckbox = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
        });

        let fakeLabel = createElement({
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Comment',
        });

        formGroup.appendChild(fakeLabel);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let formCheck = createElement({
            classList: ['custom-control', 'custom-checkbox'],
        });

        let inputAttributes = [
            {
                name: 'type',
                value: TYPES.CHECKBOX,
            },
        ];
        if (this.hasComment) {
            inputAttributes.push({
                name: 'checked',
                value: '',
            });
        }

        let input = createElement({
            type: 'input',
            classList: ['custom-control-input'],
            attributes: inputAttributes,
            id: `comment_${this.id}`,
        });

        input.addEventListener('change', () => {
            this.hasComment = input.checked;
        });

        let label = createElement({
            type: 'label',
            classList: ['custom-control-label'],
            attributes: [
                {
                    name: 'for',
                    value: `comment_${this.id}`,
                },
            ],
            innerText: 'Enable comment',
        });

        let helperText = createElement({
            type: 'small',
            classList: ['form-text', 'text-muted'],
            attributes: [
                {
                    name: 'for',
                    value: `choices_${this.id}`,
                },
            ],
            innerHTML: 'Enable this will append a textarea after the field with the option to enter an comment to their answer.',
        });

        formCheck.appendChild(input);
        formCheck.appendChild(label);
        formCheck.appendChild(helperText);
        rightSide.appendChild(formCheck);

        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Render the choices input
     * @return {Element}
     */
    this.renderChoicesTextarea = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
            id: `choices_display_${this.id}`,
        });

        if (!this.displayChoices) {
            formGroup.classList.add('sr-only');
        } else {
            formGroup.classList.remove('sr-only');
        }

        let label = createElement({
            type: 'label',
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Choices',
            attributes: [
                {
                    name: 'for',
                    value: `choices_${this.id}`,
                }
            ]
        });

        formGroup.appendChild(label);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let input = createElement({
            type: 'textarea',
            classList: ['form-control'],
            attributes: [
                {
                    name: 'placeholder',
                    value: 'Choices',
                },
            ],
            id: `choices_${this.id}`,
            value: this.choices,
        });

        input.addEventListener('keyup', () => {
            this.choices = input.value;
        });

        let helperText = createElement({
            type: 'small',
            classList: ['form-text', 'text-muted'],
            attributes: [
                {
                    name: 'for',
                    value: `choices_${this.id}`,
                },
            ],
            innerHTML: 'Enter each choice on a new line.',
        });

        rightSide.appendChild(input);
        rightSide.appendChild(helperText);
        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Render the weight input
     * @return {Element}
     */
    this.renderWeightInput = function() {
        let formGroup = createElement({
            classList: ['form-group', 'row'],
        });

        let label = createElement({
            type: 'label',
            classList: ['col-md-3', 'col-form-label'],
            innerText: 'Weight',
            attributes: [
                {
                    name: 'for',
                    value: `weight_${this.id}`,
                }
            ]
        });

        formGroup.appendChild(label);

        let rightSide = createElement({
            classList: ['col-md-9'],
        });

        let input = createElement({
            type: 'input',
            classList: ['form-control'],
            attributes: [
                {
                    name: 'type',
                    value: 'number',
                },
                {
                    name: 'placeholder',
                    value: 'Weight',
                },
                {
                    name: 'min',
                    value: 0,
                },
            ],
            id: `weight_${this.id}`,
            value: this.weight,
        });

        input.addEventListener('change', () => {
            this.weight = input.value;
            this.postRender();
        });

        let helperText = createElement({
            type: 'small',
            classList: ['form-text', 'text-muted'],
            id: `weight_helper_${this.id}`,
            innerText: '',
        });

        rightSide.appendChild(input);
        rightSide.appendChild(helperText);
        formGroup.appendChild(rightSide);

        return formGroup;
    };

    /**
     * Handles all rendering that needs to happen at the end of rendering
     */
    this.postRender = function() {
        // TODO (Svein): Fix this later, storing the total weight,
        //  for using as calculation for showing percentage for each field with weights
        // let totalWeight = localStorage.getItem('totalWeight');
    };

    /**
     * Return field data as an object
     * @return {object}
     */
    this.get = function() {
        return {
            type:           this.type,
            name:           this.name,
            description:    this.description,
            hasComment:     this.hasComment,
            label:          this.label,
            order:          this.order,
            weight:         this.weight,
            choices:        this.getChoices(),
        }
    };

    /**
     * Return choices as a single string, with values separated by a comma
     * @return {string}
     */
    this.getChoices = function() {
        let val = this.choices.split('\n');
        return val.join(',');
    };
}

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

/**
 * Small fast agile function for creating HTMLElement's
 * @param {object} args
 * @return {Element}
 */
function createElement(args) {
    let el = document.createElement(args.type || 'div');

    if (args.classList !== undefined) {
        //noinspection JSValidateTypes
        if (args.classList.length !== undefined) {
            el.classList.add(...args.classList);
        }
    }

    if (args.id !== undefined) {
        el.id = args.id;
    }

    if (args.attributes !== undefined) {
        args.attributes.forEach(a => {
            el.setAttribute(a.name, a.value);
        });
    }

    if (args.innerText !== undefined) {
        el.innerText = args.innerText;
    }

    if (args.innerHTML !== undefined) {
        el.innerHTML = args.innerHTML;
    }

    if (args.value !== undefined) {
        el.value = args.value;
    }

    return el;
}