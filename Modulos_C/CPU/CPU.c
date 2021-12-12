#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h>
#include <linux/seq_file.h>
#include <linux/hugetlb.h>
#include <linux/sched.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM SOPES 1");
MODULE_AUTHOR("Jose Martinez");

//struct sysinfo info;
struct task_struct *proceso;
struct task_struct *task_child;
struct list_head *list;
static int escribir_archivo(struct seq_file *archivo, void *v)
{
    //long total;
    long rss;
    bool esPadre;
    bool first;
    bool next = true;
    bool conhijos;

    //seq_printf(archivo, "Lab de SOPES1\n");
    seq_printf(archivo, "{\"root\": [\n");
    for_each_process(proceso)
    {
        first = true;
        if (proceso->mm)
        {
            rss = get_mm_rss(proceso->mm);
            if (next)
            {
                seq_printf(archivo, "{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"RAM\":\"%ld\",\n\"Usuario\":\"%d\",\n\"Estado\":\"%ld\",\n", proceso->comm, proceso->pid, rss, __kuid_val(proceso->real_cred->uid),proceso->state);
                /*seq_printf(archivo, "{\"Proceso\":\"%s\",\n", proceso->comm);
                seq_printf(archivo, "\"PID\":\"%d\",\n", proceso->pid);
                seq_printf(archivo, "\"RAM\":\"%ld\",\n", rss);
                seq_printf(archivo, "\"Usuario\":\"%d\",\n", proceso->real_cred->uid);*/
                next = false;
            }
            else
            {
                seq_printf(archivo, ",{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"RAM\":\"%ld\",\n\"Usuario\":\"%d\",\n\"Estado\":\"%ld\",\n", proceso->comm, proceso->pid, rss, __kuid_val(proceso->real_cred->uid),proceso->state);

                /*seq_printf(archivo, ",{\"Proceso\":\"%s\",\n", proceso->comm);
                seq_printf(archivo, "\"PID\":\"%d\",\n", proceso->pid);
                seq_printf(archivo, "\"RAM\":\"%ld\",\n", rss);
                seq_printf(archivo, "\"Usuario\":\"%d\",\n", proceso->real_cred->uid);*/
            }
            esPadre = true;
        }
        else
        {
            if (next)
            {
                seq_printf(archivo, "{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"RAM\":\"0\",\n\"Usuario\":\"%d\",\n\"Estado\":\"%ld\",\n", proceso->comm, proceso->pid,__kuid_val(proceso->real_cred->uid),proceso->state);

                /*seq_printf(archivo, "{\"Proceso\":\"%s\",\n", proceso->comm);
                seq_printf(archivo, "\"PID\":\"%d\",\n", proceso->pid);
                seq_printf(archivo, "\"RAM\":\"0\",\n");
                seq_printf(archivo, "\"Usuario\":\"%d\",\n", proceso->real_cred->uid);*/
                next = false;
            }
            else
            {
                seq_printf(archivo, ",{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"RAM\":\"0\",\n\"Usuario\":\"%d\",\n\"Estado\":\"%ld\",\n", proceso->comm, proceso->pid,__kuid_val(proceso->real_cred->uid),proceso->state);
                /*seq_printf(archivo, ",{\"Proceso\":\"%s\",\n", proceso->comm);
                seq_printf(archivo, "\"PID\":\"%d\",\n", proceso->pid);
                seq_printf(archivo, "\"RAM\":\"0\",\n");
                seq_printf(archivo, "\"Usuario\":\"%d\",\n", proceso->real_cred->uid);*/
            }
            esPadre = true;
        }
        conhijos = false;
        seq_printf(archivo, "\"hijos\":[");
        list_for_each(list, &(proceso->children))
        {
            task_child = list_entry(list, struct task_struct, sibling);

            conhijos = true;
            //rss = get_mm_rss(task_child->mm) << PAGE_SHIFT;
            if (first)
            {
                seq_printf(archivo,"{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"Estado\":\"%d\"}",task_child->comm,task_child->pid,task_child->state);
                /*seq_printf(archivo, "{\"Proceso\": \"%s\",\n", task_child->comm);
                seq_printf(archivo, "\"PID\": \"%d\"\n}", task_child->pid);*/
                first = false;
            }
            else
            {
                seq_printf(archivo,",{\"Proceso\":\"%s\",\n\"PID\":\"%d\",\n\"Estado\":\"%d\"}",task_child->comm,task_child->pid,task_child->state);
                /*seq_printf(archivo, ",{\"Proceso\": \"%s\",\n", task_child->comm);
                seq_printf(archivo, "\"PID\": \"%d\"\n}", task_child->pid);*/
            }
        }
        if (esPadre && conhijos)
        {
            seq_printf(archivo, "]\n}\n");
        }
        else if (esPadre)
        {
            seq_printf(archivo, "]\n}\n");
        }
        else
        {
            seq_printf(archivo, "},\n");
        }

        //seq_printf(archivo,"Proceso: %s (PID: %d)\n",proceso->comm,proceso->pid);
    }
    seq_printf(archivo, "]}\n");

    //seq_printf(archivo, "Guatemala Diciembre 2021\n");

    return 0;
}

static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

static struct proc_ops operaciones = {
    .proc_open = al_abrir,
    .proc_read = seq_read};

static int _insert(void)
{
    proc_create("Modulo de CPU", 0, NULL, &operaciones);
    printk(KERN_INFO "Mensaje al insertar el modulo, Lab SO\n");

    return 0;
}

static void _remove(void)
{
    remove_proc_entry("Modulo de CPU", NULL);
    printk(KERN_INFO "Mensaje al remover el modulo, Lab SO\n");
}

module_init(_insert);
module_exit(_remove);