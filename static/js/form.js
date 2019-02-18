(function() {
    const TYPES = {
        CHECKBOX: 'checkbox',
        NUMBER: 'number',
        RADIO: 'radio',
        TEXT: 'text',
        TEXTAREA: 'textarea',
        URL: 'url',
    };

    /**
     *
     * @constructor
     */
    let Form = function(selector) {
        /**
         *
         * @type {HTMLElement}
         */
        this.output = document.getElementById(selector);

        /**
         *
         * @type {Object[]}
         */
        this.fields = [];

        /**
         *
         * @param {Field} field
         */
        this.add = function(field) {
            this.fields.push(field);
        };

        /**
         *
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
            })
        };

        /**
         *
         */
        this.newField = function() {
            this.fields.push(new Field());

            this.renderAll();
        };

        /**
         *
         */
        this.renderAll = function() {
            this.output.innerHTML = '';

            this.fields.forEach((e, i) => {
                this.output.appendChild(e.render(i));
            });
        }
    };

    /**
     *
     * @constructor
     */
    let Field = function() {
        /**
         *
         * @type {string}
         */
        this.type = TYPES.TEXT;

        /**
         *
         * @type {string}
         */
        this.name = '';

        /**
         *
         * @type {string|number}
         */
        this.value = '';

        /**
         *
         * @type {string}
         */
        this.label = '';

        /**
         *
         * @type {number}
         */
        this.order = 0;

        /**
         *
         * @param {number} id
         * @returns {HTMLElement}
         */
        this.render = function(id) {
            let card, cardHeader, cardBody, h2, button, collapse, formGroup, label, select, input, option;

            card = document.createElement('div');
            card.classList.add(...['card']);

            cardHeader = document.createElement('div');
            cardHeader.classList.add(...['card-header']);
            cardHeader.id = `heading_${id}`;

            h2 = document.createElement('h2');
            h2.classList.add(...['mb-0']);

            button = document.createElement('button');
            button.classList.add(...['btn', 'btn-link']);
            button.setAttribute('type', 'button');
            button.setAttribute('data-toggle', 'collapse');
            button.setAttribute('data-target', `#collapse_${id}`);
            button.setAttribute('aria-controls', `collapse_${id}`);
            button.setAttribute('aria-expanded', 'false');
            button.innerText = `Question ${id+1}`;

            h2.appendChild(button);
            cardHeader.appendChild(h2);

            collapse = document.createElement('div');
            collapse.classList.add(...['collapse']);
            collapse.id = `collapse_${id}`;
            collapse.setAttribute('aria-labelledby', `header_${id}`);
            collapse.setAttribute('data-parent', '#formAccordion'); // TODO (Svein): Fix this

            cardBody = document.createElement('div');
            cardBody.classList.add(...['card-body']);

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
            });

            formGroup.appendChild(label);
            formGroup.appendChild(select);
            cardBody.appendChild(formGroup);
            // == INPUT TYPE END ==

            // == LABEL START ==
            formGroup = document.createElement('div');
            formGroup.classList.add(...['form-group']);

            label = document.createElement('label');
            label.setAttribute('for', `label_${id}`);
            label.innerText = 'Label';

            input = document.createElement('input');
            input.classList.add(...['form-control']);
            input.setAttribute('type', 'text');
            input.setAttribute('name', `label_${id}`);
            input.id = `label_${id}`;
            input.value = this.label;
            // EventListener to the input-field, saving the value
            input.addEventListener('keyup', e => {
                e.preventDefault();
                this.label = input.value;
            });

            formGroup.appendChild(label);
            formGroup.appendChild(input);
            cardBody.appendChild(formGroup);
            // == LABEL END ==

            // == NAME START ==
            formGroup = document.createElement('div');
            formGroup.classList.add(...['form-group']);

            label = document.createElement('label');
            label.setAttribute('for', `name_${id}`);
            label.innerText = 'Name';

            input = document.createElement('input');
            input.classList.add(...['form-control']);
            input.setAttribute('type', 'text');
            input.setAttribute('name', `name_${id}`);
            input.id = `name_${id}`;
            input.value = this.name;
            // EventListener to the input-field, saving the value
            input.addEventListener('keyup', e => {
                e.preventDefault();
                this.name = input.value;
            });

            formGroup.appendChild(label);
            formGroup.appendChild(input);
            cardBody.appendChild(formGroup);
            // == NAME END ==


            card.appendChild(cardHeader);
            card.appendChild(collapse);

            return card;
        };
    };

    let button = document.getElementById('addField');
    let form = new Form('formAccordion');

    button.addEventListener('click', () => {
        form.newField();
    });

})();