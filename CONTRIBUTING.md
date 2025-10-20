# Contributing to NamDoS Pro

Thank you for your interest in contributing to NamDoS Pro! This document provides guidelines for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Contributing Process](#contributing-process)
- [Code Style](#code-style)
- [Testing](#testing)
- [Documentation](#documentation)
- [Security](#security)
- [Legal Considerations](#legal-considerations)

## Code of Conduct

### Our Pledge

We are committed to providing a welcoming and inspiring community for all. Please read and follow our Code of Conduct:

- **Be Respectful**: Treat everyone with respect and kindness
- **Be Inclusive**: Welcome newcomers and help them learn
- **Be Constructive**: Provide constructive feedback and suggestions
- **Be Responsible**: Use the tool ethically and responsibly

### Unacceptable Behavior

- Harassment, discrimination, or offensive comments
- Trolling, insulting, or derogatory comments
- Public or private harassment
- Publishing private information without permission
- Any conduct that could be considered inappropriate

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- Basic understanding of DDoS attack vectors
- Understanding of ethical hacking principles

### Development Setup

1. **Fork the Repository**
   ```bash
   git clone https://github.com/your-username/ddos.git
   cd ddos
   ```

2. **Set Up Development Environment**
   ```bash
   # Install dependencies
   go mod download
   
   # Install development tools
   make install-tools
   ```

3. **Create a Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Contributing Process

### 1. Choose an Issue

- Look for issues labeled `good first issue` or `help wanted`
- Comment on the issue to express interest
- Wait for maintainer approval before starting

### 2. Make Changes

- Write clean, readable code
- Follow the code style guidelines
- Add tests for new functionality
- Update documentation as needed

### 3. Test Your Changes

```bash
# Run tests
make test-unit

# Run linter
make lint

# Format code
make fmt

# Build and test
make build
make test
```

### 4. Submit a Pull Request

- Push your changes to your fork
- Create a pull request with a clear description
- Link to the related issue
- Wait for review and feedback

## Code Style

### Go Style Guidelines

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` to format code
- Use `golint` for linting
- Write clear, self-documenting code

### Naming Conventions

- **Functions**: Use camelCase, be descriptive
- **Variables**: Use camelCase, be concise
- **Constants**: Use UPPER_CASE
- **Packages**: Use lowercase, single word

### Example

```go
// Good
func calculateAttackRate(requests int, duration time.Duration) float64 {
    return float64(requests) / duration.Seconds()
}

// Bad
func calc(a int, b time.Duration) float64 {
    return float64(a) / b.Seconds()
}
```

## Testing

### Test Requirements

- All new features must have tests
- Test coverage should be maintained or improved
- Tests should be fast and reliable
- Use table-driven tests when appropriate

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...

# Run specific test
go test -v -run TestSpecificFunction
```

### Test Examples

```go
func TestCalculateAttackRate(t *testing.T) {
    tests := []struct {
        name     string
        requests int
        duration time.Duration
        expected float64
    }{
        {
            name:     "normal case",
            requests: 1000,
            duration: 60 * time.Second,
            expected: 16.67,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := calculateAttackRate(tt.requests, tt.duration)
            if math.Abs(result-tt.expected) > 0.01 {
                t.Errorf("expected %v, got %v", tt.expected, result)
            }
        })
    }
}
```

## Documentation

### Code Documentation

- Document all public functions and types
- Use clear, concise comments
- Include examples when helpful
- Update documentation with code changes

### README Updates

- Update README.md for new features
- Include usage examples
- Update installation instructions
- Keep feature lists current

### API Documentation

- Document command-line options
- Include parameter descriptions
- Provide usage examples
- Explain return values

## Security

### Security Guidelines

- **Never commit secrets**: No API keys, passwords, or tokens
- **Validate inputs**: Always validate user inputs
- **Handle errors securely**: Don't expose sensitive information
- **Use secure defaults**: Choose secure default configurations

### Security Testing

- Test with various inputs
- Check for buffer overflows
- Validate error handling
- Test edge cases

### Reporting Security Issues

- **DO NOT** create public issues for security vulnerabilities
- Email security@example.com instead
- Follow responsible disclosure practices

## Legal Considerations

### Ethical Use

- This tool is for educational and authorized testing only
- Contributors must use the tool ethically
- No malicious use or unauthorized testing
- Respect target systems and their owners

### Legal Compliance

- Ensure compliance with local laws
- Respect terms of service
- Obtain proper authorization
- Use responsibly

## Pull Request Guidelines

### Before Submitting

- [ ] Code follows style guidelines
- [ ] Tests pass and coverage is maintained
- [ ] Documentation is updated
- [ ] No security vulnerabilities introduced
- [ ] Legal and ethical considerations addressed

### Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass
- [ ] New tests added
- [ ] Manual testing completed

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] No breaking changes
```

## Review Process

### What We Look For

- **Code Quality**: Clean, readable, maintainable code
- **Functionality**: Does it work as intended?
- **Testing**: Adequate test coverage
- **Documentation**: Clear documentation
- **Security**: No security vulnerabilities
- **Performance**: Efficient implementation

### Review Timeline

- **Initial Review**: Within 3-5 business days
- **Follow-up Reviews**: Within 1-2 business days
- **Final Approval**: Within 1 week

## Getting Help

### Resources

- **Documentation**: Check README and code comments
- **Issues**: Search existing issues
- **Discussions**: Use GitHub Discussions
- **Community**: Join our community

### Contact

- **Maintainers**: @naem021
- **Email**: support@example.com
- **Discord**: [Join our Discord](https://discord.gg/example)

## Recognition

### Contributors

- All contributors are recognized in CONTRIBUTORS.md
- Significant contributions get special recognition
- Contributors are listed in release notes

### Types of Contributions

- **Code**: Bug fixes, new features, improvements
- **Documentation**: README updates, code comments
- **Testing**: Test cases, bug reports
- **Community**: Helping others, answering questions

## License

By contributing to NamDoS Pro, you agree that your contributions will be licensed under the MIT License.

---

**Thank you for contributing to NamDoS Pro! Together, we can make it better for everyone.**
