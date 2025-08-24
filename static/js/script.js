// static/js/script.js
document.addEventListener('DOMContentLoaded', function() {
    console.log('Gin app loaded successfully!');
    
    // Example: Add click event to buttons
    const buttons = document.querySelectorAll('.btn');
    buttons.forEach(btn => {
        btn.addEventListener('click', function() {
            alert('Button clicked!');
        });
    });
});