using Microsoft.EntityFrameworkCore;
using TodoList.Models;

namespace TodoList.Data
{
    public class TodoDbContext : DbContext
    {
        public TodoDbContext() { }
        public TodoDbContext(DbContextOptions<TodoDbContext> options) : base(options) { }

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            base.OnConfiguring(optionsBuilder);
            optionsBuilder.UseNpgsql("User ID=lqxmrbbg;Password=nkRjn98SGwuybPjnZ9ipAEaD3jcZSJG1;Host=horton.db.elephantsql.com;Port=5432;Database=lqxmrbbg;Pooling=true;Connection Lifetime=0");
        }

        public DbSet<Todo> Todos { get; set; }
    }
}
