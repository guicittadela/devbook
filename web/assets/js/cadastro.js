$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault()
    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas nao coincidem")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        alert("Usuario cadastrado")
    }).fail(function(erro){
        console.log(erro)
        alert("erro")
    });
    return
}