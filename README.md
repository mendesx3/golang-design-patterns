# Princípios SOLID em Go

Aqui estão os cinco princípios SOLID da programação orientada a objetos aplicados ao desenvolvimento em Go:

### 1. SRP - Princípio da Responsabilidade Única

- **Objetivo**: Uma classe deve ter apenas um motivo para mudar.
- **Benefícios**: Manutenção simplificada, melhor encapsulamento.

### 2. OCP - Princípio do Aberto/Fechado

- **Objetivo**: Entidades de software (classes, módulos, funções) devem estar abertas para extensão, mas fechadas para modificação.
- **Benefícios**: Facilita a extensão de código sem afetar o código existente.

### 3. LSP - Princípio da Substituição de Liskov

- **Objetivo**: Objetos de uma classe derivada devem ser substituíveis pelos objetos da classe base sem afetar a integridade do programa.
- **Benefícios**: Polimorfismo seguro, código mais robusto.

### 4. ISP - Princípio da Segregação de Interfaces

- **Objetivo**: Muitas interfaces específicas são melhores do que uma única interface geral.
- **Benefícios**: Evita interfaces inchadas, classes cliente não precisam implementar métodos que não usam.

### 5. DIP - Princípio da Inversão de Dependência

- **Objetivo**: Módulos de alto nível não devem depender de módulos de baixo nível; ambos devem depender de abstrações. Abstrações não devem depender de detalhes; detalhes devem depender de abstrações.
- **Benefícios**: Desacoplamento, facilidade de teste e manutenção.


# Padrões de Projeto em Go

Aqui estão alguns dos principais padrões de projeto em Go:

### 1. Singleton

- **Objetivo**: Garante uma única instância de uma classe e acesso global a ela.
- **Uso Comum**: Gerenciadores de configuração, registros de log.

### 2. Factory

- **Objetivo**: Define uma interface para criar objetos, permitindo que as subclasses escolham a classe concreta.
- **Uso Comum**: Bibliotecas de criação de widgets gráficos, abstração de criação de recursos.

### 3. Builder

- **Objetivo**: Separa a construção de objetos complexos de sua representação, permitindo várias construções.
- **Uso Comum**: Construção de consultas SQL complexas, construtores de documentos.

### 4. Prototype

- **Objetivo**: Cria novos objetos a partir de objetos existentes, útil quando a criação direta é complexa.
- **Uso Comum**: Clonagem de objetos, geração de objetos de teste.

### 5. Adapter

- **Objetivo**: Permite que a interface de uma classe seja usada como uma interface diferente.
- **Uso Comum**: Adaptação de bibliotecas legadas, conectores de API.

### 6. Decorator

- **Objetivo**: Anexa responsabilidades adicionais a objetos dinamicamente.
- **Uso Comum**: Adicionar funcionalidade a objetos sem modificar suas classes.

### 7. Observer

- **Objetivo**: Define uma dependência um-para-muitos entre objetos, notifica mudanças.
- **Uso Comum**: Tratamento de eventos, atualização de interfaces de usuário.

### 8. Strategy

- **Objetivo**: Define uma família de algoritmos, tornando-os intercambiáveis.
- **Uso Comum**: Seleção de algoritmo em tempo de execução, implementação de diferentes estratégias de ordenação.

### 9. Command

- **Objetivo**: Encapsula uma solicitação como um objeto, permitindo parametrização, fila e execução adiada.
- **Uso Comum**: Filas de comandos, suporte para desfazer e refazer.

### 10. State

- **Objetivo**: Permite que um objeto altere seu comportamento quando seu estado interno muda.
- **Uso Comum**: Máquinas de estados, gerenciamento de comportamento condicional.

### 11. Chain of Responsibility

- **Objetivo**: Passa solicitações ao longo de uma cadeia de manipuladores até que uma seja processada.
- **Uso Comum**: Tratamento de solicitações HTTP, processamento de filtros.

### 12. Template Method

- **Objetivo**: Define o esqueleto de um algoritmo, permitindo que as subclasses redefinam etapas específicas.
- **Uso Comum**: Implementação de algoritmos comuns com variações específicas.

### 13. Composite

- **Objetivo**: Compose objetos em uma estrutura de árvore para tratar objetos individuais e composições de forma uniforme.
- **Uso Comum**: Modelagem de árvores, estruturas de menu.

### 14. Visitor

- **Objetivo**: Representa uma operação a ser executada nos elementos de uma estrutura de objeto.
- **Uso Comum**: Análise de árvores, processamento de elementos em uma coleção.

### 15. Proxy

- **Objetivo**: Fornece um substituto ou espaço reservado para controlar o acesso a um objeto.
- **Uso Comum**: Acesso a recursos remotos, controle de acesso a objetos caros.

