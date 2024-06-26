document.getElementById('api-form').addEventListener('submit', function(event) {
    event.preventDefault();
    const apiKey = document.getElementById('api_key').value;
    const companyUrl = document.getElementById('company_url').value;
    const currentYear = new Date().getFullYear();
    const currentMonth = String(new Date().getMonth() + 1).padStart(2, '0'); // 月は0から始まるため+1
    document.cookie = `token=${apiKey}; path=/`;
    document.cookie = `company_url=${companyUrl}; path=/`;
    window.location.href = `/aggregate?monthly=${currentYear}-${currentMonth}`;
});
