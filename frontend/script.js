document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('ascii-form');

    if (form) {
        form.addEventListener('submit', function (event) {
            event.preventDefault();

            const formData = new FormData(form);

            fetch('/ascii-art', {
                method: 'POST',
                body: formData
            })
                .then(response => {
                    if (response.ok) {
                        return response.text();
                    }
                    console.log(response);
                    throw new Error('Network response was not ok.');
                })
                .then(data => {
                    console.log('Success:', data);
                    const outputBox = document.getElementById('output-box');
                    if (outputBox) {
                        outputBox.textContent = data;
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
        });
    }
});
