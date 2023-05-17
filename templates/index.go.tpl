{{ define "index" }} {{ template "Header" }}

<table class="table">
  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">NAME</th>
      <th scope="col">EMAIL</th>
      <th scope="col">ACTIONS</th>
    </tr>
  </thead>
  <tbody>
    {{ range.}}
    <tr>
      <th scope="row">{{.ID}}</th>
      <td>{{.Name}}</td>
      <td>{{.Email}}</td>
      <td>
        <button class="btn btn-danger">DELETE</button>
        <button class="btn btn-primary">UPDATE</button>
      </td>
    </tr>
    {{ end }}
  </tbody>
</table>

{{ template "Footer" }} {{ end }}
